package notes

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// SearchRequest represents an Search request.
type SearchRequest struct {
	Query   string      `json:"query"`
	Limit   uint        `json:"limit"`
	SinceID string      `json:"sinceId"`
	UntilID string      `json:"untilId"`
	Host    core.String `json:"host"`
	UserID  core.String `json:"userId"`
}

// Validate the request.
func (r SearchRequest) Validate() error {
	if r.Query == "" {
		return core.RequestValidationError{
			Request: r,
			Message: core.UndefinedRequiredField,
			Field:   "Query",
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

// Search endpoint.
func (s *Service) Search(request SearchRequest) ([]models.Note, error) {
	var response []models.Note

	err := s.Call(
		&core.JSONRequest{Request: &request, Path: "/notes/search"},
		&response,
	)

	return response, err
}
