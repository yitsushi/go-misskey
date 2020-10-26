package notes

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// RenotesRequest represents an Renotes request.
type RenotesRequest struct {
	NoteID  string `json:"noteId"`
	Limit   uint   `json:"limit"`
	SinceID string `json:"sinceId"`
	UntilID string `json:"untilId"`
}

// Validate the request.
func (r RenotesRequest) Validate() error {
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

// Renotes endpoint.
func (s *Service) Renotes(request RenotesRequest) ([]models.Note, error) {
	var response []models.Note

	err := s.Call(
		&core.JSONRequest{Request: &request, Path: "/notes/renotes"},
		&response,
	)

	return response, err
}
