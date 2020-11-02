package promo

import (
	"github.com/yitsushi/go-misskey/core"
)

// MarkAsReadRequest represents an MarkAsRead request.
type MarkAsReadRequest struct {
	NoteID string `json:"noteId"`
}

// Validate the request.
func (r MarkAsReadRequest) Validate() error {
	return nil
}

// MarkAsRead endpoint.
func (s *Service) MarkAsRead(noteID string) error {
	request := MarkAsReadRequest{NoteID: noteID}

	err := s.Call(
		&core.JSONRequest{Request: &request, Path: "/promo/read"},
		&core.EmptyResponse{},
	)

	return err
}
