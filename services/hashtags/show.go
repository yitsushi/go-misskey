package hashtags

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// ShowRequest represents an Show request.
type ShowRequest struct {
	Tag string `json:"tag"`
}

// Validate the request.
func (r ShowRequest) Validate() error {
	return nil
}

// Show endpoint.
func (s *Service) Show(tag string) (models.Hashtag, error) {
	var response models.Hashtag

	request := ShowRequest{Tag: tag}

	err := s.Call(
		&core.JSONRequest{Request: &request, Path: "/hashtags/show"},
		&response,
	)

	return response, err
}
