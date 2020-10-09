package meta

import (
	"github.com/yitsushi/go-misskey/core"
)

// AnnouncementsRequest represents an Announcement request.
type AnnouncementsRequest struct {
	*core.BaseRequest
	WithUnreads bool   `json:"withUnreads"`
	SinceID     string `json:"sinceId"`
	UntilID     string `json:"untilId"`
}

// AnnouncementsResponse represents an Announcement response.
type AnnouncementsResponse []Announcement

// AnnouncementOptions is the options list for Announcement().
type AnnouncementOptions struct {
	WithUnreads bool
	SinceID     string
	UntilID     string
}

// Announcements lists all announcements.
func (s *Service) Announcements(options *AnnouncementOptions) (AnnouncementsResponse, error) {
	request := &AnnouncementsRequest{
		WithUnreads: options.WithUnreads,
		SinceID:     options.SinceID,
		UntilID:     options.UntilID,
	}

	var response AnnouncementsResponse
	s.Call(
		&core.BaseRequest{Request: request, Path: "/announcements"},
		&response,
	)

	return response, nil
}

