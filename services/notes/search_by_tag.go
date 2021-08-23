package notes

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// SearchByTagRequest represents an SearchByTag request.
type SearchByTagRequest struct {
	Tag           string      `json:"tag,omitempty"`
	Query         [][]string  `json:"query,omitempty"`
	Limit         uint        `json:"limit"`
	Reply         bool        `json:"reply,omitempty"`
	Renote        bool        `json:"renote,omitempty"`
	OnlyWithFiles bool        `json:"withFiles,omitempty"`
	OnlyPolls     bool        `json:"poll,omitempty"`
	SinceID       string      `json:"sinceId,omitempty"`
	UntilID       string      `json:"untilId,omitempty"`
	Host          core.String `json:"host,omitempty"`
	UserID        core.String `json:"userId,omitempty"`
}

// Validate the request.
func (r SearchByTagRequest) Validate() error {
	if r.Tag == "" && len(r.Query) == 0 {
		return core.RequestValidationError{
			Request: r,
			Message: core.UndefinedRequiredField,
			Field:   "Tag || Query",
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

// SearchByTag endpoint.
func (s *Service) SearchByTag(request SearchByTagRequest) ([]models.Note, error) {
	var response []models.Note

	err := s.Call(
		&core.JSONRequest{Request: &request, Path: "/notes/search-by-tag"},
		&response,
	)

	return response, err
}
