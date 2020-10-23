package folders

import (
	"github.com/yitsushi/go-misskey/core"
)

// DeleteRequest is the request to delete a folder.
type DeleteRequest struct {
	FolderID string `json:"folderId"`
}

// Delete a folder.
func (s *Service) Delete(folderID string) error {
	request := &DeleteRequest{
		FolderID: folderID,
	}

	var response core.DummyResponse
	err := s.Call(
		&core.BaseRequest{Request: request, Path: "/drive/folders/delete"},
		&response,
	)

	return err
}
