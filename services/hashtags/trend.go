package hashtags

import (
	"github.com/yitsushi/go-misskey/core"
)

// TrendRequest represents an Trend request.
type TrendRequest struct {
}

// Trend represents one item in the response of the Trend response.
type Trend struct {
	Tag        string   `json:"tag"`
	Chart      []uint64 `json:"chart"`
	UsersCount uint64   `json:"usersCount"`
}

// Validate the request.
func (r TrendRequest) Validate() error {
	return nil
}

// Trend endpoint.
func (s *Service) Trend() ([]Trend, error) {
	var response []Trend

	err := s.Call(
		&core.JSONRequest{Request: &TrendRequest{}, Path: "/hashtags/trend"},
		&response,
	)

	return response, err
}
