package antennas

import "github.com/yitsushi/go-misskey/core"

// DeleteRequest represents a request to delete an antenna.
type DeleteRequest struct {
	AntennaID string `json:"antennaId"`
}

// Validate the request.
func (r DeleteRequest) Validate() error {
	return nil
}

// Delete is the endpoint to delete an existing antenna.
func (s *Service) Delete(antennaID string) error {
	var response core.EmptyResponse
	err := s.Call(
		&core.JSONRequest{
			Request: &DeleteRequest{
				AntennaID: antennaID,
			},
			Path: "/antennas/delete",
		},
		&response,
	)

	return err
}
