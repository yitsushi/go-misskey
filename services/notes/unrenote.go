package notes

import (
	"github.com/yitsushi/go-misskey/core"
)

// UnrenoteRequest represents an Unrenote request.
type UnrenoteRequest struct {
	NoteID string `json:"noteId"`
}

// Validate the request.
func (r UnrenoteRequest) Validate() error {
	return nil
}

// Unrenote endpoint.
func (s *Service) Unrenote(noteID string) error {
	err := s.Call(
		&core.JSONRequest{
			Request: &UnrenoteRequest{
				NoteID: noteID,
			},
			Path: "/notes/unrenote",
		},
		&core.EmptyResponse{},
	)

	return err
}
