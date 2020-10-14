package antennas

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// ShowRequest is the request object for a Show request.
type ShowRequest struct {
	AntennaID string `json:"antennaId"`
}

// Show is the endpoint to get an Antenna.
func (s *Service) Show(antennaID string) (models.Antenna, error) {
	request := &ShowRequest{
		AntennaID: antennaID,
	}

	var response models.Antenna
	err := s.Call(
		&core.BaseRequest{Request: request, Path: "/antennas/show"},
		&response,
	)

	return response, err
}
