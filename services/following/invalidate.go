package following

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// InvalidateRequest is the request structure to invalidate a following.
type InvalidateRequest struct {
	UserID string `json:"userId"`
}

// Validate request.
func (r InvalidateRequest) Validate() error {
	if r.UserID == "" {
		return core.RequestValidationError{
			Request: r,
			Message: core.UndefinedRequiredField,
			Field:   "UserID",
		}
	}

	return nil
}

// Invalidate following endpoint.
func (s *Service) Invalidate(userID string) (models.User, error) {
	var response models.User

	request := &InvalidateRequest{UserID: userID}

	err := s.Call(
		&core.JSONRequest{Request: request, Path: "/following/invalidate"},
		&response,
	)

	return response, err
}
