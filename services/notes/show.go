package notes

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// ShowRequest represents an Show request.
type ShowRequest struct {
	NoteID string `json:"noteId"`
}

// Validate the request.
func (r ShowRequest) Validate() error {
	return nil
}

// Show endpoint.
func (s *Service) Show(noteID string) (models.Note, error) {
	var response models.Note

	request := ShowRequest{NoteID: noteID}

	err := s.Call(
		&core.JSONRequest{Request: &request, Path: "/notes/show"},
		&response,
	)

	return response, err
}
