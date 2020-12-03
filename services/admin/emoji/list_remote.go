package emoji

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// ListRemoteRequest represents a ListRemote request.
type ListRemoteRequest struct {
	Query   string `json:"query,omitempty"`
	Host    string `json:"host,omitempty"`
	Limit   uint   `json:"limit,omitempty"`
	SinceID string `json:"sinceId,omitempty"`
	UntilID string `json:"untilId,omitempty"`
}

// Validate the request.
func (r ListRemoteRequest) Validate() error {
	return nil
}

// ListRemote lists all remote emojies.
func (s *Service) ListRemote(request ListRemoteRequest) ([]models.Emoji, error) {
	var response []models.Emoji

	err := s.Call(
		&core.JSONRequest{Request: &request, Path: "/admin/emoji/list-remote"},
		&response,
	)

	return response, err
}
