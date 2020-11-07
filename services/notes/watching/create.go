package watching

import (
	"github.com/yitsushi/go-misskey/core"
)

// CreateRequest represents an Create request.
type CreateRequest struct {
	NoteID string `json:"noteId"`
}

// Validate the request.
func (r CreateRequest) Validate() error {
	return nil
}

// Create endpoint.
func (s *Service) Create(noteID string) error {
	return s.Call(
		&core.JSONRequest{
			Request: &CreateRequest{
				NoteID: noteID,
			},
			Path: "/notes/watching/create",
		},
		&core.EmptyResponse{},
	)
}
