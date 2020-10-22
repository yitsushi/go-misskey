package files

import (
	"github.com/yitsushi/go-misskey/core"
)

// DeleteRequest is the request to delete a file.
type DeleteRequest struct {
	FileID string `json:"fileId"`
}

// Delete a file.
func (s *Service) Delete(fileID string) error {
	request := &DeleteRequest{
		FileID: fileID,
	}

	var response core.DummyResponse
	err := s.Call(
		&core.BaseRequest{Request: request, Path: "/drive/files/delete"},
		&response,
	)

	return err
}
