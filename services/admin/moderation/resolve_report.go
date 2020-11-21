package moderation

import (
	"github.com/yitsushi/go-misskey/core"
)

// ResolveReportRequest represents an ResolveReport Announcement request.
type ResolveReportRequest struct {
	ID string `json:"reportId"`
}

// Validate the request.
func (r ResolveReportRequest) Validate() error {
	if r.ID == "" {
		return core.RequestValidationError{
			Request: r,
			Message: core.UndefinedRequiredField,
			Field:   "ID",
		}
	}

	return nil
}

// ResolveReport lists all announcements.
func (s *Service) ResolveReport(id string) error {
	err := s.Call(
		&core.JSONRequest{
			Request: &ResolveReportRequest{ID: id},
			Path:    "/admin/resolve-abuse-user-report",
		},
		&core.EmptyResponse{},
	)

	return err
}
