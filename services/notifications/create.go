package notifications

import (
	"github.com/yitsushi/go-misskey/core"
)

// CreateRequest represents an Create request.
type CreateRequest struct {
	Body   string      `json:"body"`
	Header core.String `json:"header"`
	Icon   core.String `json:"icon"`
}

// Validate the request.
func (r CreateRequest) Validate() error {
	return nil
}

// Create endpoint.
func (s *Service) Create(request CreateRequest) error {
	var response core.DummyResponse
	err := s.Call(
		&core.JSONRequest{Request: &request, Path: "/notifications/create"},
		&response,
	)

	return err
}
