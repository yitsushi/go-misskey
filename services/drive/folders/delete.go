package folders

import (
	"github.com/yitsushi/go-misskey/core"
)

// DeleteRequest is the request to delete a folder.
type DeleteRequest struct {
	FolderID string `json:"folderId"`
}

// Validate the request.
func (r DeleteRequest) Validate() error {
	return nil
}

// Delete folder.
func (s *Service) Delete(folderID string) error {
	return s.Call(
		&core.JSONRequest{
			Request: &DeleteRequest{
				FolderID: folderID,
			},
			Path: "/drive/folders/delete",
		},
		&core.EmptyResponse{},
	)
}
