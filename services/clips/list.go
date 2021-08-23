package clips

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// ListRequest represents an List request.
type ListRequest struct{}

// Validate the request.
func (r ListRequest) Validate() error {
	return nil
}

// List clips.
func (s *Service) List() ([]models.Clip, error) {
	var response []models.Clip
	err := s.Call(
		&core.JSONRequest{Request: &ListRequest{}, Path: "/clips/list"},
		&response,
	)

	return response, err
}
