package core

import (
	"encoding/json"
)

// JSONRequest is the base request.
type JSONRequest struct {
	Path    string
	Request BaseRequest
}

const jsonContentType = "application/json"

// Validate the request.
func (r JSONRequest) Validate() error {
	return r.Request.Validate()
}

// EndpointPath returns with the path for the endpoint.
func (r JSONRequest) EndpointPath() string {
	return r.Path
}

// ToBody returns with the JSON []byte representation of the request.
func (r JSONRequest) ToBody(token string) ([]byte, string, error) {
	requestBody, _ := json.Marshal(r.Request)

	var repack map[string]interface{}

	if err := json.Unmarshal(requestBody, &repack); err != nil {
		return requestBody, jsonContentType, err
	}

	repack["i"] = token

	content, err := json.Marshal(repack)

	return content, jsonContentType, err
}
