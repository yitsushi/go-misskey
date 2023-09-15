package users

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

type MeRequest struct{}

func (r MeRequest) Validate() error {
	return nil
}

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
