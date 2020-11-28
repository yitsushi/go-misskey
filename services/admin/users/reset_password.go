package users

import (
	"github.com/yitsushi/go-misskey/core"
)

// ResetPasswordRequest is the request object for a ResetPassword request.
type ResetPasswordRequest struct {
	UserID string `json:"userId"`
}

// Validate the request.
func (r ResetPasswordRequest) Validate() error {
	if r.UserID == "" {
		return core.RequestValidationError{
			Request: r,
			Message: core.UndefinedRequiredField,
			Field:   "UserID",
		}
	}

	return nil
}

type resetPasswordResponse struct {
	Password string `json:"password"`
}

// ResetPassword is the endpoint to get a User.
func (s *Service) ResetPassword(userID string) (string, error) {
	var response resetPasswordResponse

	err := s.Call(
		&core.JSONRequest{
			Request: &ResetPasswordRequest{
				UserID: userID,
			},
			Path: "/admin/reset-password",
		},
		&response,
	)
	if err != nil {
		return "", err
	}

	return response.Password, nil
}
