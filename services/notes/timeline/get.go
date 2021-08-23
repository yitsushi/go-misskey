package timeline

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// GetRequest represents an Get request.
type GetRequest struct {
	Limit                 uint   `json:"limit"`
	SinceID               string `json:"sinceId"`
	UntilID               string `json:"untilId"`
	SinceDate             uint64 `json:"sinceDate"`
	UntilDate             uint64 `json:"untilDate"`
	IncludeMyRenotes      bool   `json:"includeMyRenotes"`
	IncludeRenotedMyNotes bool   `json:"includeRenotedMyNotes"`
	IncludeLocalRenotes   bool   `json:"includeLocalRenotes"`
	OnlyWithFiles         bool   `json:"withFiles"`
}

// Validate the request.
func (r GetRequest) Validate() error {
	if r.Limit < 1 || r.Limit > maxLimit {
		return core.RequestValidationError{
			Request: r,
			Message: core.NewRangeError(1, maxLimit),
			Field:   "Limit",
		}
	}

	return nil
}

// Get is the endpoint for /notes/timeline.
func (s *Service) Get(request GetRequest) ([]models.Note, error) {
	var response []models.Note

	err := s.Call(
		&core.JSONRequest{Request: &request, Path: "/notes/timeline"},
		&response,
	)

	return response, err
}
