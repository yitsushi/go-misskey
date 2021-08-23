package notes

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// FeaturedRequest represents an Featured request.
type FeaturedRequest struct {
	Limit  uint   `json:"limit"`
	Offset uint64 `json:"offset"`
}

// Validate the request.
func (r FeaturedRequest) Validate() error {
	if r.Limit < 1 || r.Limit > maxLimit {
		return core.RequestValidationError{
			Request: r,
			Message: core.NewRangeError(1, maxLimit),
			Field:   "Limit",
		}
	}

	return nil
}

// Featured endpoint.
func (s *Service) Featured(request FeaturedRequest) ([]models.Note, error) {
	var response []models.Note

	err := s.Call(
		&core.JSONRequest{Request: &request, Path: "/notes/featured"},
		&response,
	)

	return response, err
}
