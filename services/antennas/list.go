package antennas

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// ListRequest is empty, but keep it as it is for consistency.
type ListRequest struct {
}

// Validate the request.
func (r ListRequest) Validate() error {
	return nil
}

// List is the endpoint to list all existing antennas.
func (s *Service) List() ([]models.Antenna, error) {
	var response []models.Antenna
	err := s.Call(
		&core.JSONRequest{Request: &ListRequest{}, Path: "/antennas/list"},
		&response,
	)

	return response, err
}
