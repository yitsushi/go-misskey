package notes

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// MentionsRequest represents an Mentions request.
type MentionsRequest struct {
	Following  string            `json:"following,omitempty"`
	Limit      uint              `json:"limit"`
	SinceID    string            `json:"sinceId,omitempty"`
	UntilID    string            `json:"untilId,omitempty"`
	Visibility models.Visibility `json:"visibility,omitempty"`
}

// Validate the request.
func (r MentionsRequest) Validate() error {
	if r.Limit < 1 || r.Limit > maxLimit {
		return core.RequestValidationError{
			Request: r,
			Message: core.NewRangeError(1, maxLimit),
			Field:   "Limit",
		}
	}

	return nil
}

// Mentions endpoint.
func (s *Service) Mentions(request MentionsRequest) ([]models.Note, error) {
	var response []models.Note

	err := s.Call(
		&core.JSONRequest{Request: &request, Path: "/notes/mentions"},
		&response,
	)

	return response, err
}
