package federation

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// ShowInstanceRequest contains request information to obtain a single instance.
type ShowInstanceRequest struct {
	Host string `json:"host"`
}

// Validate the request.
func (r ShowInstanceRequest) Validate() error {
	if r.Host == "" {
		return core.RequestValidationError{
			Request: r,
			Message: core.UndefinedRequiredField,
			Field:   "Host",
		}
	}

	return nil
}

// ShowInstance returns a single instance given a host.
func (s *Service) ShowInstance(request ShowInstanceRequest) (models.Instance, error) {
	var response models.Instance

	err := s.Call(
		&core.JSONRequest{Request: &request, Path: "/federation/show-instance"},
		&response,
	)

	return response, err
}
