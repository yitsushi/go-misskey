package relays

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// ListRequest represents a List relays request.
type ListRequest struct{}

// Validate the request.
func (r ListRequest) Validate() error {
	return nil
}

// List a relay.
func (s *Service) List() ([]models.Relay, error) {
	response := []models.Relay{}
	err := s.Call(
		&core.JSONRequest{Request: &ListRequest{}, Path: "/admin/relays/list"},
		&response,
	)

	return response, err
}
