package announcements

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// ListRequest represents an List Announcement request.
type ListRequest struct {
	Limit   uint   `json:"limit"`
	SinceID string `json:"sinceId,omitempty"`
	UntilID string `json:"untilId,omitempty"`
}

// Validate the request.
func (r ListRequest) Validate() error {
	if r.Limit < 1 || r.Limit > 100 {
		return core.RequestValidationError{
			Request: r,
			Message: core.NewRangeError(1, 100),
			Field:   "Limit",
		}
	}

	return nil
}

// List lists all announcements.
func (s *Service) List(request ListRequest) ([]models.Announcement, error) {
	var response []models.Announcement
	err := s.Call(
		&core.JSONRequest{Request: &request, Path: "/admin/announcements/list"},
		&response,
	)

	return response, err
}
