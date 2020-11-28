package users

import (
	"github.com/yitsushi/go-misskey/core"
)

// SilenceRequest is the request object for a Silence request.
type SilenceRequest struct {
	UserID string `json:"userId"`
}

// Validate the request.
func (r SilenceRequest) Validate() error {
	if r.UserID == "" {
		return core.RequestValidationError{
			Request: r,
			Message: core.UndefinedRequiredField,
			Field:   "UserID",
		}
	}

	return nil
}

// Silence is the endpoint to get a User.
func (s *Service) Silence(userID string) error {
	return s.Call(
		&core.JSONRequest{
			Request: &SilenceRequest{
				UserID: userID,
			},
			Path: "/admin/silence-user",
		},
		&core.EmptyResponse{},
	)
}
