package antennas

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// ShowRequest is the request object for a Show request.
type ShowRequest struct {
	AntennaID string `json:"antennaId"`
}

// Validate the request.
func (r ShowRequest) Validate() error {
	return nil
}

// Show is the endpoint to get an Antenna.
func (s *Service) Show(antennaID string) (models.Antenna, error) {
	var response models.Antenna
	err := s.Call(
		&core.JSONRequest{
			Request: &ShowRequest{
				AntennaID: antennaID,
			},
			Path: "/antennas/show",
		},
		&response,
	)

	return response, err
}
