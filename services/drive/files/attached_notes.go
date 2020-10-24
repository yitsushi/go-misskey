package files

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// AttachedNotesRequest list all notes where a given file has reference.
type AttachedNotesRequest struct {
	FileID string `json:"fileId"`
}

// Validate the request.
func (r AttachedNotesRequest) Validate() error {
	return nil
}

// AttachedNotes gets drive information.
func (s *Service) AttachedNotes(fileID string) ([]models.Note, error) {
	var response []models.Note
	err := s.Call(
		&core.JSONRequest{
			Request: &AttachedNotesRequest{
				FileID: fileID,
			},
			Path: "/drive/files/attached-notes",
		},
		&response,
	)

	return response, err
}
