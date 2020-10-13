package antennas

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/entities"
)

// ListRequest is empty, but keep it as it is for consistency.
type ListRequest struct {
}

// ListResponse is an array of Antennas from the list response.
type ListResponse []entities.Antenna

// List is the endpoint to list all existing antennas.
func (s *Service) List() (ListResponse, error) {
	request := &ListRequest{}

	var response ListResponse
	err := s.Call(
		&core.BaseRequest{Request: request, Path: "/antennas/list"},
		&response,
	)

	return response, err
}
