package hashtags

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// ShowRequest represents an Show request.
type ShowRequest struct {
	Tag string `json:"tag"`
}

// Show endpoint.
func (s *Service) Show(tag string) (models.Hashtag, error) {
	var response models.Hashtag

	request := ShowRequest{Tag: tag}

	err := s.Call(
		&core.BaseRequest{Request: &request, Path: "/hashtags/show"},
		&response,
	)

	return response, err
}
