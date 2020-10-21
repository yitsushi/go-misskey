package drive

import (
	"github.com/yitsushi/go-misskey/core"
)

// InformationRequest is a pseudo type to get information about drive.
// No payload, but that way requests stay consistent.
type InformationRequest struct{}

type InformationResponse struct {
	Capacity core.DataSize `json:"capacity"`
	Usage    core.DataSize `json:"usage"`
}

// Information gets drive information.
func (s *Service) Information() (InformationResponse, error) {
	request := &InformationRequest{}

	var response InformationResponse
	err := s.Call(
		&core.BaseRequest{Request: request, Path: "/drive"},
		&response,
	)

	return response, err
}
