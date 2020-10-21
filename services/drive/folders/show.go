package folders

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// ShowRequest is the request to show a folder.
type ShowRequest struct {
	FolderID string `json:"folderId"`
}

// Show gets a folder by its ID.
func (s *Service) Show(folderID string) (models.Folder, error) {
	request := &ShowRequest{
		FolderID: folderID,
	}

	var response models.Folder
	err := s.Call(
		&core.BaseRequest{Request: request, Path: "/drive/folders/show"},
		&response,
	)

	return response, err
}
