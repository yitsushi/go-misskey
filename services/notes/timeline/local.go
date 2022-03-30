package timeline

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// LocalRequest represents an Local request.
type LocalRequest struct {
	OnlyWithFileType []string `json:"fileType,omitempty"`
	OnlyWithFiles    bool     `json:"withFiles"`
	ExcludeNSFW      bool     `json:"excludeNsfw"`
	Limit            uint     `json:"limit"`
	SinceID          string   `json:"sinceId,omitempty"`
	UntilID          string   `json:"untilId,omitempty"`
	SinceDate        uint64   `json:"sinceDate,omitempty"`
	UntilDate        uint64   `json:"untilDate,omitempty"`
}

// Validate the request.
func (r LocalRequest) Validate() error {
	if r.Limit < 1 || r.Limit > maxLimit {
		return core.RequestValidationError{
			Request: r,
			Message: core.NewRangeError(1, maxLimit),
			Field:   "Limit",
		}
	}

	return nil
}

// Local is the endpoint for /notes/local-timeline.
func (s *Service) Local(request LocalRequest) ([]models.Note, error) {
	var response []models.Note

	err := s.Call(
		&core.JSONRequest{Request: &request, Path: "/notes/local-timeline"},
		&response,
	)

	return response, err
}
