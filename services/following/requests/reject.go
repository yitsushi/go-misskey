package requests

import (
	"github.com/yitsushi/go-misskey/core"
)

// RejectRequest is the request structure to reject a following request.
type RejectRequest struct {
	UserID string `json:"userId"`
}

// Validate request.
func (r RejectRequest) Validate() error {
	if r.UserID == "" {
		return core.RequestValidationError{
			Request: r,
			Message: core.UndefinedRequiredField,
			Field:   "UserID",
		}
	}

	return nil
}

// Reject following request endpoint.
func (s *Service) Reject(userID string) error {
	request := &RejectRequest{UserID: userID}

	err := s.Call(
		&core.JSONRequest{Request: request, Path: "/following/requests/reject"},
		&core.EmptyResponse{},
	)

	return err
}
