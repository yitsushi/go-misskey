package users

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// CreateRequest is the request object for a Create request.
type CreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Validate the request.
func (r CreateRequest) Validate() error {
	if r.Username == "" {
		return core.RequestValidationError{
			Request: r,
			Message: core.UndefinedRequiredField,
			Field:   "Username",
		}
	}

	if r.Password == "" {
		return core.RequestValidationError{
			Request: r,
			Message: core.UndefinedRequiredField,
			Field:   "Password",
		}
	}

	return nil
}

// Create is the endpoint to get a User.
func (s *Service) Create(username string, password string) (models.UserFromAdmin, error) {
	var response models.UserFromAdmin

	err := s.Call(
		&core.JSONRequest{
			Request: &CreateRequest{
				Username: username,
				Password: password,
			},
			Path: "/admin/accounts/create",
		},
		&response,
	)

	return response, err
}
