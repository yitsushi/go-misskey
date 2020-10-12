package antennas

import "github.com/yitsushi/go-misskey/core"

// DeleteRequest represents a request to delete an antenna.
type DeleteRequest struct {
	AntennaID string `json:"antennaId"`
}

// DeleteResponse represents a response to delete an antenna.
// Even if it's empty, for consistency, there is a response type for it.
type DeleteResponse struct {
}

// Delete is the endpoint to delete an existing antenna.
func (s *Service) Delete(antennaID string) (DeleteResponse, error) {
	request := &DeleteRequest{
		AntennaID: antennaID,
	}

	var response DeleteResponse
	err := s.Call(
		&core.BaseRequest{Request: request, Path: "/antennas/delete"},
		&response,
	)

	return response, err
}
