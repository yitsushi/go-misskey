package files

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// AttachedNotesRequest list all notes where a given file has reference.
type AttachedNotesRequest struct {
	FileID string `json:"fileId"`
}

// AttachedNotes gets drive information.
func (s *Service) AttachedNotes(fileID string) ([]models.Note, error) {
	request := &AttachedNotesRequest{
		FileID: fileID,
	}

	var response []models.Note
	err := s.Call(
		&core.BaseRequest{Request: request, Path: "/drive/files/attached-notes"},
		&response,
	)

	return response, err
}
