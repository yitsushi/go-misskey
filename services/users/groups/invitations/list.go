package invitations

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// ListRequest represents an List request.
type ListRequest struct {
	Limit   int    `json:"limit,omitempty"`
	SinceID string `json:"sinceId,omitempty"`
	UntilID string `json:"untilId,omitempty"`
}

// Validate the request.
func (r ListRequest) Validate() error {
	return nil
}

// List group.
func (s *Service) List(request ListRequest) ([]models.GroupInvitation, error) {
	response := []models.GroupInvitation{}
	err := s.Call(
		&core.JSONRequest{Request: &request, Path: "/i/user-group-invites"},
		&response,
	)

	return response, err
}
