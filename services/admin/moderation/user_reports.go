package moderation

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// UserReportsRequest represents an UserReports Announcement request.
type UserReportsRequest struct {
	Limit            uint        `json:"limit,omitempty"`
	SinceID          string      `json:"sinceId,omitempty"`
	UntilID          string      `json:"untilId,omitempty"`
	State            ReportState `json:"state,omitempty"`
	ReporterOrigin   UserOrigin  `json:"reporterOrigin,omitempty"`
	TargetUserOrigin UserOrigin  `json:"targetUserOrigin,omitempty"`
}

// Validate the request.
func (r UserReportsRequest) Validate() error {
	return nil
}

// UserReports lists all announcements.
func (s *Service) UserReports(request UserReportsRequest) ([]models.Report, error) {
	var response []models.Report
	err := s.Call(
		&core.JSONRequest{Request: &request, Path: "/admin/abuse-user-reports"},
		&response,
	)

	return response, err
}
