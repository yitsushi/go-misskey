package core

import (
	"bytes"
	"mime/multipart"
)

// MultipartRequest is the base multipart form request.
type MultipartRequest struct {
	Path    string
	Request BaseRequest
}

// Validate the request.
func (r MultipartRequest) Validate() error {
	return r.Request.Validate()
}

// EndpointPath returns with the path for the endpoint.
func (r MultipartRequest) EndpointPath() string {
	return r.Path
}

// ToBody generates the multipart request body.
func (r *MultipartRequest) ToBody(token string) ([]byte, string, error) {
	body := bytes.Buffer{}
	writer := multipart.NewWriter(&body)

	tokenPart, _ := writer.CreateFormField("i")

	_, err := tokenPart.Write([]byte(token))
	if err != nil {
		writer.Close()

		return body.Bytes(), "", err
	}

	fields := parseMultipartFields(r.Request)
	for _, field := range fields {
		if field.Ref != "" {
			if _, ok := fields[field.Ref]; !ok {
				return body.Bytes(), "", InvalidFieldReferenceError{
					Name:      field.Name,
					Type:      field.Type,
					Reference: field.Ref,
				}
			}

			field.Name = fields[field.Ref].Name
		}

		writeField(*writer, field)
	}

	writer.Close()

	return body.Bytes(), writer.FormDataContentType(), nil
}

func writeField(writer multipart.Writer, field multipartField) {
	switch field.Type {
	case "field":
		part, _ := writer.CreateFormField(field.Name)
		_, _ = part.Write(field.Value)
	case "file":
		part, _ := writer.CreateFormFile("file", field.Name)
		_, _ = part.Write(field.Value)
	}
}
