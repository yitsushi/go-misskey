package core

import (
	"encoding/json"
)

// BaseRequest is the base request.
type BaseRequest struct {
	Path    string
	Request interface{}
}

const jsonContentType = "application/json"

// EndpointPath returns with the path for the endpoint.
func (r BaseRequest) EndpointPath() string {
	return r.Path
}

// ToBody returns with the JSON []byte representation of the request.
func (r BaseRequest) ToBody(token string) ([]byte, string, error) {
	requestBody, _ := json.Marshal(r.Request)

	var repack map[string]interface{}

	if err := json.Unmarshal(requestBody, &repack); err != nil {
		return requestBody, jsonContentType, err
	}

	repack["i"] = token

	content, err := json.Marshal(repack)

	return content, jsonContentType, err
}
