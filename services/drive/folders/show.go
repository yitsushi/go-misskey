package folders

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// ShowRequest is the request to show a folder.
type ShowRequest struct {
	FolderID string `json:"folderId"`
}

// Validate the request.
func (r ShowRequest) Validate() error {
	return nil
}

// Show gets a folder by its ID.
func (s *Service) Show(folderID string) (models.Folder, error) {
	var response models.Folder
	err := s.Call(
		&core.JSONRequest{
			Request: &ShowRequest{
				FolderID: folderID,
			},
			Path: "/drive/folders/show",
		},
		&response,
	)

	return response, err
}
