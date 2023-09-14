package users

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

func (s *Service) Me() (models.User, error) {
	var user models.User
	err := s.Call(
		&core.JSONRequest{
			Path: "/i",
		},
		&user,
	)

	return user, err
}
