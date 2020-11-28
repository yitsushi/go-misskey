package users

import (
	"github.com/yitsushi/go-misskey/core"
)

// UnsuspendRequest is the request object for a Unsuspend request.
type UnsuspendRequest struct {
	UserID string `json:"userId"`
}

// Validate the request.
func (r UnsuspendRequest) Validate() error {
	if r.UserID == "" {
		return core.RequestValidationError{
			Request: r,
			Message: core.UndefinedRequiredField,
			Field:   "UserID",
		}
	}

	return nil
}

// Unsuspend is the endpoint to get a User.
func (s *Service) Unsuspend(userID string) error {
	return s.Call(
		&core.JSONRequest{
			Request: &UnsuspendRequest{
				UserID: userID,
			},
			Path: "/admin/unsuspend-user",
		},
		&core.EmptyResponse{},
	)
}
