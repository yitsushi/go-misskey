package notes

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// ConversationRequest represents an Conversation request.
type ConversationRequest struct {
	NoteID string `json:"noteId"`
	Limit  uint   `json:"limit"`
	Offset uint64 `json:"offset,omitempty"`
}

// Validate the request.
func (r ConversationRequest) Validate() error {
	if r.NoteID == "" {
		return core.RequestValidationError{
			Request: r,
			Message: core.UndefinedRequiredField,
			Field:   "NoteID",
		}
	}

	if r.Limit < 1 || r.Limit > 100 {
		return core.RequestValidationError{
			Request: r,
			Message: core.NewRangeError(1, 100),
			Field:   "Limit",
		}
	}

	return nil
}

// Conversation endpoint.
func (s *Service) Conversation(request ConversationRequest) ([]models.Note, error) {
	var response []models.Note

	err := s.Call(
		&core.JSONRequest{Request: &request, Path: "/notes/conversation"},
		&response,
	)

	return response, err
}
