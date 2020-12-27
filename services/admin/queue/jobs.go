package queue

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// JobsRequest represents an Jobs request.
type JobsRequest struct {
	Domain Domain `json:"domain"`
	State  State  `json:"state"`
	Limit  uint   `json:"limit,omitempty"`
}

// Validate the request.
func (r JobsRequest) Validate() error {
	if r.Domain == "" {
		return core.RequestValidationError{
			Request: r,
			Message: core.UndefinedRequiredField,
			Field:   "Domain",
		}
	}

	if r.State == "" {
		return core.RequestValidationError{
			Request: r,
			Message: core.UndefinedRequiredField,
			Field:   "State",
		}
	}

	return nil
}

// Jobs lists all jobs in a queue with state.
func (s *Service) Jobs(request JobsRequest) ([]models.Job, error) {
	var response []models.Job

	err := s.Call(
		&core.JSONRequest{Request: &request, Path: "/admin/queue/jobs"},
		&response,
	)

	return response, err
}
