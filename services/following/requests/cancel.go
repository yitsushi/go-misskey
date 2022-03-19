package requests

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// CancelRequest is the request structure to cancel a following request.
type CancelRequest struct {
	UserID string `json:"userId"`
}

// Validate request.
func (r CancelRequest) Validate() error {
	if r.UserID == "" {
		return core.RequestValidationError{
			Request: r,
			Message: core.UndefinedRequiredField,
			Field:   "UserID",
		}
	}

	return nil
}

// Cancel following endpoint.
func (s *Service) Cancel(userID string) (models.User, error) {
	var response models.User

	request := &CancelRequest{UserID: userID}

	err := s.Call(
		&core.JSONRequest{Request: request, Path: "/following/requests/cancel"},
		&response,
	)

	return response, err
}
