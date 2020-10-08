package core

import (
	"encoding/json"
)

// BaseRequest is the base request.
type BaseRequest struct {
	APIToken string
	Path     string
	Request  interface{}
}

// ToJSON returns with the JSON []byte representation of the request.
func (r BaseRequest) ToJSON() ([]byte, error) {
	requestBody, _ := json.Marshal(r.Request)

	var repack map[string]interface{}

	if err := json.Unmarshal(requestBody, &repack); err != nil {
		return requestBody, err
	}

	repack["i"] = r.APIToken

	return json.Marshal(repack)
}

// SetAPIToken stores the API key for a Misskey User.
func (r *BaseRequest) SetAPIToken(token string) {
	r.APIToken = token
}
