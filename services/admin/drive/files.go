package drive

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// FilesRequest represents a Files request.
type FilesRequest struct {
	Limit    uint              `json:"limit,omitempty"`
	SinceID  string            `json:"sinceId,omitempty"`
	UntilID  string            `json:"untilId,omitempty"`
	Type     core.String       `json:"type,omitempty"`
	Origin   models.UserOrigin `json:"origin,omitempty"`
	Hostname core.String       `json:"hostname,omitempty"`
}

// Validate the request.
func (r FilesRequest) Validate() error {
	return nil
}

// Files lists all emojies.
func (s *Service) Files(request FilesRequest) ([]models.File, error) {
	var response []models.File

	err := s.Call(
		&core.JSONRequest{Request: &request, Path: "/admin/drive/files"},
		&response,
	)

	return response, err
}
