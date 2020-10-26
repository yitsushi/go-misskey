package notes

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// ReactionsRequest represents an Reactions request.
type ReactionsRequest struct {
	NoteID  string      `json:"noteId"`
	Type    core.String `json:"type"`
	Limit   uint        `json:"Limit"`
	Offset  uint        `json:"offset"`
	SinceID string      `json:"sinceId"`
	UntilID string      `json:"untilId"`
}

// Validate the request.
func (r ReactionsRequest) Validate() error {
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

// Reactions endpoint.
func (s *Service) Reactions(request ReactionsRequest) ([]models.Reaction, error) {
	var response []models.Reaction

	err := s.Call(
		&core.JSONRequest{Request: &request, Path: "/notes/reactions"},
		&response,
	)

	return response, err
}
