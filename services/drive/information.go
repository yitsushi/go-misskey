package drive

import (
	"github.com/yitsushi/go-misskey/core"
)

// InformationRequest is a pseudo type to get information about drive.
// No payload, but that way requests stay consistent.
type InformationRequest struct{}

// InformationResponse is the representation of the /drive/files request.
type InformationResponse struct {
	Capacity core.DataSize `json:"capacity"`
	Usage    core.DataSize `json:"usage"`
}

// Validate the request.
func (r InformationRequest) Validate() error {
	return nil
}

// Information gets drive information.
func (s *Service) Information() (InformationResponse, error) {
	var response InformationResponse
	err := s.Call(
		&core.JSONRequest{Request: &InformationRequest{}, Path: "/drive"},
		&response,
	)

	return response, err
}
