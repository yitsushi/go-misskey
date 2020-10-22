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

// AnnouncementOptions is the options list for Announcement().
type AnnouncementOptions struct {
	WithUnreads bool
	SinceID     string
	UntilID     string
}

// Announcements lists all announcements.
func (s *Service) Announcements(options *AnnouncementOptions) ([]models.Announcement, error) {
	request := &AnnouncementsRequest{
		WithUnreads: options.WithUnreads,
		SinceID:     options.SinceID,
		UntilID:     options.UntilID,
	}

	var response []models.Announcement
	err := s.Call(
		&core.BaseRequest{Request: request, Path: "/announcements"},
		&response,
	)

	return response, err
}
