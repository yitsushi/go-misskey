package core_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/go-misskey/core"
)

type ExampleMultipartRequest struct {
	Name    string `multipart:"name,type=field"`
	Content []byte `multipart:"ref=name,type=file"`
}

// Validate the request.
func (r ExampleMultipartRequest) Validate() error {
	return nil
}

func TestMultipartRequest(t *testing.T) {
	request := core.MultipartRequest{
		Path: "/test/endpoint",
		Request: ExampleMultipartRequest{
			Name:    "filename",
			Content: []byte("File Content"),
		},
	}

	body, contentType, err := request.ToBody("api-token")
	if !assert.NoError(t, err) {
		return
	}

	stringBody := string(body)

	assert.Contains(t, contentType, "multipart/form-data; boundary=")
	assert.Contains(t, stringBody, `Content-Disposition: form-data; name="i"`)
	assert.Contains(t, stringBody, `Content-Disposition: form-data; name="name"`)
	assert.Contains(t, stringBody, `Content-Disposition: form-data; name="file"; filename="name"`)
	assert.Contains(t, stringBody, `Content-Type: application/octet-stream`)
	assert.Contains(t, stringBody, `api-token`)
	assert.Contains(t, stringBody, `File Content`)
}
