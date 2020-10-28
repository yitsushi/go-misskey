package files

import (
	"github.com/yitsushi/go-misskey/core"
)

// DeleteRequest is the request to delete a file.
type DeleteRequest struct {
	FileID string `json:"fileId"`
}

// Validate the request.
func (r DeleteRequest) Validate() error {
	return nil
}

// Delete a file.
func (s *Service) Delete(fileID string) error {
	return s.Call(
		&core.JSONRequest{
			Request: &DeleteRequest{
				FileID: fileID,
			},
			Path: "/drive/files/delete",
		},
		&core.EmptyResponse{},
	)
}
