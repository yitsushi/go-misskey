package users

import (
	"github.com/yitsushi/go-misskey/core"
)

// SuspendRequest is the request object for a Suspend request.
type SuspendRequest struct {
	UserID string `json:"userId"`
}

// Validate the request.
func (r SuspendRequest) Validate() error {
	if r.UserID == "" {
		return core.RequestValidationError{
			Request: r,
			Message: core.UndefinedRequiredField,
			Field:   "UserID",
		}
	}

	return nil
}

// Suspend is the endpoint to get a User.
func (s *Service) Suspend(userID string) error {
	return s.Call(
		&core.JSONRequest{
			Request: &SuspendRequest{
				UserID: userID,
			},
			Path: "/admin/suspend-user",
		},
		&core.EmptyResponse{},
	)
}
