package groups

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// OwnedRequest represents an List request.
type OwnedRequest struct{}

// Validate the request.
func (r OwnedRequest) Validate() error {
	return nil
}

// Owned clips.
func (s *Service) Owned() ([]models.Group, error) {
	var response []models.Group
	err := s.Call(
		&core.JSONRequest{Request: &OwnedRequest{}, Path: "/users/groups/owned"},
		&response,
	)

	return response, err
}
