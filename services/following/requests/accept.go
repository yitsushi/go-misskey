package requests

import (
	"github.com/yitsushi/go-misskey/core"
)

// AcceptRequest is the request structure to accept a following request.
type AcceptRequest struct {
	UserID string `json:"userId"`
}

// Validate request.
func (r AcceptRequest) Validate() error {
	if r.UserID == "" {
		return core.RequestValidationError{
			Request: r,
			Message: core.UndefinedRequiredField,
			Field:   "UserID",
		}
	}

	return nil
}

// Accept following endpoint.
func (s *Service) Accept(userID string) error {
	request := &AcceptRequest{UserID: userID}

	err := s.Call(
		&core.JSONRequest{Request: request, Path: "/following/requests/accept"},
		&core.EmptyResponse{},
	)

	return err
}
