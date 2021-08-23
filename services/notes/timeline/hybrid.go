package timeline

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// HybridRequest represents an Hybrid request.
type HybridRequest struct {
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
func (r HybridRequest) Validate() error {
	if r.Limit < 1 || r.Limit > maxLimit {
		return core.RequestValidationError{
			Request: r,
			Message: core.NewRangeError(1, maxLimit),
			Field:   "Limit",
		}
	}

	return nil
}

// Hybrid is the endpoint for /notes/hybrid-timeline.
func (s *Service) Hybrid(request HybridRequest) ([]models.Note, error) {
	var response []models.Note

	err := s.Call(
		&core.JSONRequest{Request: &request, Path: "/notes/hybrid-timeline"},
		&response,
	)

	return response, err
}
