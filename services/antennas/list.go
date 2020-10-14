package antennas

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// ListRequest is empty, but keep it as it is for consistency.
type ListRequest struct {
}

// List is the endpoint to list all existing antennas.
func (s *Service) List() ([]models.Antenna, error) {
	request := &ListRequest{}

	var response []models.Antenna
	err := s.Call(
		&core.BaseRequest{Request: request, Path: "/antennas/list"},
		&response,
	)

	return response, err
}
