package drive

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// ShowRequest represents a Show request.
type ShowRequest struct {
	FileID string `json:"fileId,omitempty"`
	URL    string `json:"url,omitempty"`
}

// Validate the request.
func (r ShowRequest) Validate() error {
	return nil
}

// Show a file.
func (s *Service) Show(request ShowRequest) (models.File, error) {
	var response models.File

	err := s.Call(
		&core.JSONRequest{Request: &request, Path: "/admin/drive/show-file"},
		&response,
	)

	return response, err
}
