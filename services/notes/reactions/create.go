package reactions

import (
	"github.com/yitsushi/go-misskey/core"
)

// CreateRequest represents an /reactions request.
type CreateRequest struct {
	NoteID   string `json:"noteId"`
	Reaction string `json:"reaction"`
}

// Validate the request.
func (r CreateRequest) Validate() error {
	if r.NoteID == "" {
		return core.RequestValidationError{
			Request: r,
			Message: core.UndefinedRequiredField,
			Field:   "NoteID",
		}
	}

	if r.Reaction == "" {
		return core.RequestValidationError{
			Request: r,
			Message: core.UndefinedRequiredField,
			Field:   "Reaction",
		}
	}

	return nil
}

// Create endpoint.
func (s *Service) Create(request CreateRequest) error {
	err := s.Call(
		&core.JSONRequest{Request: &request, Path: "/notes/reactions/create"},
		&core.EmptyResponse{},
	)

	return err
}
