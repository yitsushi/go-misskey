package timeline

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// GlobalRequest represents an Global request.
type GlobalRequest struct {
	OnlyWithFiles bool   `json:"withFiles"`
	Limit         uint   `json:"limit"`
	SinceID       string `json:"sinceId,omitempty"`
	UntilID       string `json:"untilId,omitempty"`
	SinceDate     uint64 `json:"sinceDate,omitempty"`
	UntilDate     uint64 `json:"untilDate,omitempty"`
}

// Validate the request.
func (r GlobalRequest) Validate() error {
	if r.Limit < 1 || r.Limit > maxLimit {
		return core.RequestValidationError{
			Request: r,
			Message: core.NewRangeError(1, maxLimit),
			Field:   "Limit",
		}
	}

	return nil
}

// Global is the endpoint for /notes/global-timeline.
func (s *Service) Global(request GlobalRequest) ([]models.Note, error) {
	var response []models.Note

	err := s.Call(
		&core.JSONRequest{Request: &request, Path: "/notes/global-timeline"},
		&response,
	)

	return response, err
}
