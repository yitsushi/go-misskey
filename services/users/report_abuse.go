package users

import (
	"github.com/yitsushi/go-misskey/core"
)

// ReportAbuseRequest represents a ReportAbuse request.
type ReportAbuseRequest struct {
	UserID  string `json:"userId"`  // required
	Comment string `json:"comment"` // required
}

// Validate the request.
func (r ReportAbuseRequest) Validate() error {
	if r.UserID == "" {
		return core.RequestValidationError{
			Request: r,
			Message: core.UndefinedRequiredField,
			Field:   "UserID",
		}
	}

	if r.Comment == "" {
		return core.RequestValidationError{
			Request: r,
			Message: core.UndefinedRequiredField,
			Field:   "Comment",
		}
	}

	return nil
}

// ReportAbuse is used to file a report.
// It requires authentication.
func (s *Service) ReportAbuse(userID string, comment string) error {
	request := ReportAbuseRequest{UserID: userID, Comment: comment}

	err := s.Call(
		&core.JSONRequest{Request: &request, Path: "/users/report-abuse"},
		&core.EmptyResponse{},
	)

	return err
}
