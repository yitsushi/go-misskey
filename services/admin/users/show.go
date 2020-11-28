package users

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// ShowRequest is the request object for a Show request.
type ShowRequest struct {
	UserID string `json:"userId"`
}

// Validate the request.
func (r ShowRequest) Validate() error {
	if r.UserID == "" {
		return core.RequestValidationError{
			Request: r,
			Message: core.UndefinedRequiredField,
			Field:   "UserID",
		}
	}

	return nil
}

// Show is the endpoint to get a User.
func (s *Service) Show(userID string) (models.UserFromAdmin, error) {
	var response models.UserFromAdmin

	err := s.Call(
		&core.JSONRequest{
			Request: &ShowRequest{
				UserID: userID,
			},
			Path: "/admin/show-user",
		},
		&response,
	)

	return response, err
}
