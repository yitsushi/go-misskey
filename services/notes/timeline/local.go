package timeline

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// LocalRequest represents an Local request.
type LocalRequest struct {
	OnlyWithFiles    bool     `json:"withFiles"`
	OnlyWithFileType []string `json:"fileType,omitempty"`
	ExcludeNSFW      bool     `json:"excludeNsfw"`
	Limit            uint     `json:"limit"`
	SinceID          string   `json:"sinceId"`
	UntilID          string   `json:"untilId"`
	SinceDate        uint64   `json:"sinceDate"`
	UntilDate        uint64   `json:"untilDate"`
}

// Validate the request.
func (r LocalRequest) Validate() error {
	if r.Limit < 1 || r.Limit > 100 {
		return core.RequestValidationError{
			Request: r,
			Message: core.NewRangeError(1, 100),
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
