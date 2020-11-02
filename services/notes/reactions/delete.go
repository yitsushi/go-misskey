package reactions

import (
	"github.com/yitsushi/go-misskey/core"
)

// DeleteRequest represents an /reactions request.
type DeleteRequest struct {
	NoteID string `json:"noteId"`
}

// Validate the request.
func (r DeleteRequest) Validate() error {
	if r.NoteID == "" {
		return core.RequestValidationError{
			Request: r,
			Message: core.UndefinedRequiredField,
			Field:   "NoteID",
		}
	}

	return nil
}

// Delete endpoint.
func (s *Service) Delete(noteID string) error {
	err := s.Call(
		&core.JSONRequest{Request: &DeleteRequest{NoteID: noteID}, Path: "/notes/reactions/delete"},
		&core.EmptyResponse{},
	)

	return err
}
