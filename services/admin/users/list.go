package users

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// ListRequest is the request object for a List request.
type ListRequest struct {
	Limit    uint              `json:"limit,omitempty"`
	Offset   uint64            `json:"offset,omitempty"`
	Sort     string            `json:"sort,omitempty"`
	State    UserState         `json:"state,omitempty"`
	Origin   models.UserOrigin `json:"origin,omitempty"`
	Username string            `json:"username,omitempty"`
	Hostname string            `json:"hostname,omitempty"`
}

// Validate the request.
func (r ListRequest) Validate() error {
	return nil
}

// List is the endpoint to get a User.
func (s *Service) List(request ListRequest) ([]models.User, error) {
	var response []models.User

	err := s.Call(
		&core.JSONRequest{
			Request: &request,
			Path:    "/admin/show-users",
		},
		&response,
	)

	return response, err
}
