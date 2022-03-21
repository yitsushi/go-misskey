package groups

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// CreateRequest represents an Create request.
type CreateRequest struct {
	Name string `json:"name"`
}

// Validate the request.
func (r CreateRequest) Validate() error {
	if r.Name == "" {
		return core.RequestValidationError{
			Request: r,
			Message: core.UndefinedRequiredField,
			Field:   "Name",
		}
	}

	return nil
}

// Create group.
func (s *Service) Create(name string) (models.Group, error) {
	request := CreateRequest{Name: name}
	response := models.Group{}
	err := s.Call(
		&core.JSONRequest{Request: &request, Path: "/users/groups/create"},
		&response,
	)

	return response, err
}
