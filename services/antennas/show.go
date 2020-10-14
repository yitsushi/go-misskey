package antennas

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// ShowRequest is the request object for a Show request.
type ShowRequest struct {
	AntennaID string `json:"antennaId"`
}

// ShowResponse is the response object for a Show request.
type ShowResponse models.Antenna

// Antenna returns a single Antenna resource from the response.
func (r *ShowResponse) Antenna() models.Antenna {
	return models.Antenna(*r)
}

// Show is the endpoint to get an Antenna.
func (s *Service) Show(antennaID string) (ShowResponse, error) {
	request := &ShowRequest{
		AntennaID: antennaID,
	}

	var response ShowResponse
	err := s.Call(
		&core.BaseRequest{Request: request, Path: "/antennas/show"},
		&response,
	)

	return response, err
}
