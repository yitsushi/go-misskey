package files

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// FindRequest is the request to find file(s) by their and parent folder.
type FindRequest struct {
	Name     string      `json:"name"`
	FolderID core.String `json:"folderId"`
}

// Validate the request.
func (r FindRequest) Validate() error {
	return nil
}

// Find gets file(s) by their name and parent folder.
func (s *Service) Find(request FindRequest) ([]models.File, error) {
	var response []models.File
	err := s.Call(
		&core.JSONRequest{Request: &request, Path: "/drive/files/find"},
		&response,
	)

	return response, err
}
