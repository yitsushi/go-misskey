package notes

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// ChildrenRequest represents an Children request.
type ChildrenRequest struct {
	NoteID  string `json:"noteId"`
	Limit   uint   `json:"limit"`
	SinceID string `json:"sinceId,omitempty"`
	UntilID string `json:"untilId,omitempty"`
}

// Validate the request.
func (r ChildrenRequest) Validate() error {
	if r.NoteID == "" {
		return core.RequestValidationError{
			Request: r,
			Message: core.UndefinedRequiredField,
			Field:   "NoteID",
		}
	}

	if r.Limit < 1 || r.Limit > maxLimit {
		return core.RequestValidationError{
			Request: r,
			Message: core.NewRangeError(1, maxLimit),
			Field:   "Limit",
		}
	}

	return nil
}

// Children endpoint.
func (s *Service) Children(request ChildrenRequest) ([]models.Note, error) {
	var response []models.Note

	err := s.Call(
		&core.JSONRequest{Request: &request, Path: "/notes/children"},
		&response,
	)

	return response, err
}
