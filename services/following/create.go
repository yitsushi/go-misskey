package following

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// CreateRequest is the request structure to create a following.
type CreateRequest struct {
	UserID string `json:"userId"`
}

// Validate request.
func (r CreateRequest) Validate() error {
	if r.UserID == "" {
		return core.RequestValidationError{
			Request: r,
			Message: core.UndefinedRequiredField,
			Field:   "UserID",
		}
	}

	return nil
}

// Create following endpoint.
func (s *Service) Create(userID string) (models.User, error) {
	var response models.User

	request := &CreateRequest{UserID: userID}

	err := s.Call(
		&core.JSONRequest{Request: request, Path: "/following/create"},
		&response,
	)

	return response, err
}
