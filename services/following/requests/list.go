package requests

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// ListRequest is the request structure to list following requests.
type ListRequest struct{}

// Validate request.
func (r ListRequest) Validate() error {
	return nil
}

// ListResponse is the response from the list following requests endpoint.
type ListResponse struct {
	ID       string      `json:"id"`
	Follower models.User `json:"follower"`
	Followee models.User `json:"followee"`
}

// List following requests endpoint.
func (s *Service) List() ([]ListResponse, error) {
	response := []ListResponse{}

	err := s.Call(
		&core.JSONRequest{Request: &ListRequest{}, Path: "/following/requests/list"},
		&response,
	)

	return response, err
}
