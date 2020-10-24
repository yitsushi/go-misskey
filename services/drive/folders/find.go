package folders

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// FindRequest is the request to find folder(s) by their and parent folder.
type FindRequest struct {
	Name     string      `json:"name"`
	ParentID core.String `json:"parentId"`
}

// Validate the request.
func (r FindRequest) Validate() error {
	return nil
}

// Find gets folder(s) by their name and parent folder.
func (s *Service) Find(request FindRequest) ([]models.Folder, error) {
	var response []models.Folder
	err := s.Call(
		&core.JSONRequest{Request: &request, Path: "/drive/folders/find"},
		&response,
	)

	return response, err
}
