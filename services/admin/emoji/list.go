package emoji

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// ListRequest represents a List request.
type ListRequest struct {
	Query   string `json:"query,omitempty"`
	Limit   uint   `json:"limit,omitempty"`
	SinceID string `json:"sinceId,omitempty"`
	UntilID string `json:"untilId,omitempty"`
}

// Validate the request.
func (r ListRequest) Validate() error {
	return nil
}

// List lists all emojies.
func (s *Service) List(request ListRequest) ([]models.Emoji, error) {
	var response []models.Emoji

	err := s.Call(
		&core.JSONRequest{Request: &request, Path: "/admin/emoji/list"},
		&response,
	)

	return response, err
}
