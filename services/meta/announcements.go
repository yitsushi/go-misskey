package meta

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// AnnouncementsRequest represents an Announcement request.
type AnnouncementsRequest struct {
	WithUnreads bool   `json:"withUnreads"`
	SinceID     string `json:"sinceId"`
	UntilID     string `json:"untilId"`
}

// Validate the request.
func (r AnnouncementsRequest) Validate() error {
	return nil
}

// Announcements lists all announcements.
func (s *Service) Announcements(request AnnouncementsRequest) ([]models.Announcement, error) {
	var response []models.Announcement
	err := s.Call(
		&core.JSONRequest{Request: &request, Path: "/announcements"},
		&response,
	)

	return response, err
}
