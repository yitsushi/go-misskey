package favorites

import (
	"github.com/yitsushi/go-misskey/core"
)

// DeleteRequest represents an Delete request.
type DeleteRequest struct {
	NoteID string `json:"noteId"`
}

// Validate the request.
func (r DeleteRequest) Validate() error {
	return nil
}

// Delete endpoint.
func (s *Service) Delete(noteID string) error {
	return s.Call(
		&core.JSONRequest{
			Request: &DeleteRequest{
				NoteID: noteID,
			},
			Path: "/notes/favorites/delete",
		},
		&core.EmptyResponse{},
	)
}
