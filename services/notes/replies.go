package notes

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// RepliesRequest represents an Replies request.
type RepliesRequest struct {
	NoteID  string `json:"noteId"`
	Limit   uint   `json:"limit"`
	SinceID string `json:"sinceId"`
	UntilID string `json:"untilId"`
}

// Validate the request.
func (r RepliesRequest) Validate() error {
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

// Replies endpoint.
func (s *Service) Replies(request RepliesRequest) ([]models.Note, error) {
	var response []models.Note

	err := s.Call(
		&core.JSONRequest{Request: &request, Path: "/notes/replies"},
		&response,
	)

	return response, err
}
