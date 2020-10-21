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

type FindOptions struct {
	Name     string
	FolderID core.String
}

// Find gets file(s) by their name and parent folder.
func (s *Service) Find(options *FindOptions) ([]models.File, error) {
	request := &FindRequest{
		Name:     options.Name,
		FolderID: options.FolderID,
	}

	var response []models.File
	err := s.Call(
		&core.BaseRequest{Request: request, Path: "/drive/files/find"},
		&response,
	)

	return response, err
}
