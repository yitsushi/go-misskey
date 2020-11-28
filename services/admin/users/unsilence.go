package users

import (
	"github.com/yitsushi/go-misskey/core"
)

// UnsilenceRequest is the request object for a Unsilence request.
type UnsilenceRequest struct {
	UserID string `json:"userId"`
}

// Validate the request.
func (r UnsilenceRequest) Validate() error {
	if r.UserID == "" {
		return core.RequestValidationError{
			Request: r,
			Message: core.UndefinedRequiredField,
			Field:   "UserID",
		}
	}

	return nil
}

// Unsilence is the endpoint to get a User.
func (s *Service) Unsilence(userID string) error {
	return s.Call(
		&core.JSONRequest{
			Request: &UnsilenceRequest{
				UserID: userID,
			},
			Path: "/admin/unsilence-user",
		},
		&core.EmptyResponse{},
	)
}
