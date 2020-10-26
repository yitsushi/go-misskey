package notes

import (
	"github.com/yitsushi/go-misskey/core"
)

// StateRequest represents an State request.
type StateRequest struct {
	NoteID string `json:"noteId"`
}

// Validate the request.
func (r StateRequest) Validate() error {
	return nil
}

// NoteState is the state of a note.
type NoteState struct {
	IsFavorited bool `json:"isFavorited"`
	IsWatching  bool `json:"isWatching"`
}

// State endpoint.
func (s *Service) State(noteID string) (NoteState, error) {
	var response NoteState

	request := StateRequest{NoteID: noteID}

	err := s.Call(
		&core.JSONRequest{Request: &request, Path: "/notes/state"},
		&response,
	)

	return response, err
}
