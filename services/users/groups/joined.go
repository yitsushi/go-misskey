package groups

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// JoinedRequest represents an List request.
type JoinedRequest struct{}

// Validate the request.
func (r JoinedRequest) Validate() error {
	return nil
}

// Joined clips.
func (s *Service) Joined() ([]models.Group, error) {
	var response []models.Group
	err := s.Call(
		&core.JSONRequest{Request: &JoinedRequest{}, Path: "/users/groups/joined"},
		&response,
	)

	return response, err
}
