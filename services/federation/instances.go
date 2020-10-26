package federation

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// InstancesRequest contains request information to obtain instances.
type InstancesRequest struct {
	Host          string `json:"host"`
	Blocked       bool   `json:"blocked"`
	NotResponding bool   `json:"notResponding"`
	Suspended     bool   `json:"suspended"`
	Federating    bool   `json:"federating"`
	Subscribing   bool   `json:"subscribing"`
	Publishing    bool   `json:"publishing"`
	Limit         int    `json:"limit"`
	Offset        int    `json:"offset"`
	Sort          string `json:"sort"`
}

// Validate the request.
func (r *InstancesRequest) Validate() error {
	if r.Host == "" {
		return core.RequestValidationError{
			Request: r,
			Message: core.UndefinedRequiredField,
			Field:   "Host",
		}
	}

	if r.Limit < 1 || r.Limit > 100 {
		return core.RequestValidationError{
			Request: r,
			Message: core.NewRangeError(1, 100),
			Field:   "Limit",
		}
	}

	return nil
}

// Instances will list all instances and information about the instances for a given host.
func (s *Service) Instances(request InstancesRequest) ([]models.Instance, error) {
	var response []models.Instance

	err := s.Call(
		&core.JSONRequest{Request: &request, Path: "/federation/instances"},
		&response,
	)

	return response, err
}
