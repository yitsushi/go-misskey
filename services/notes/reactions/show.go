package reactions

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// ShowRequest represents an /reactions request.
type ShowRequest struct {
	NoteID  string      `json:"noteId"`
	Type    core.String `json:"type"`
	Limit   uint        `json:"Limit"`
	Offset  uint        `json:"offset"`
	SinceID string      `json:"sinceId"`
	UntilID string      `json:"untilId"`
}

// Validate the request.
func (r ShowRequest) Validate() error {
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

// Show endpoint.
func (s *Service) Show(request ShowRequest) ([]models.Reaction, error) {
	var response []models.Reaction

	err := s.Call(
		&core.JSONRequest{Request: &request, Path: "/notes/reactions"},
		&response,
	)

	return response, err
}
