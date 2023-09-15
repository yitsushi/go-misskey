package users

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// MeRequest represents a request to retrieve a user who has been issued a token.
type MeRequest struct{}

// Validate the request.
func (r MeRequest) Validate() error {
	return nil
}

// Me gets the user issued the token.
func (s *Service) Me() (models.User, error) {
	var user models.User
	err := s.Call(
		&core.JSONRequest{
			Path:    "/i",
			Request: &MeRequest{},
		},
		&user,
	)

	return user, err
}
