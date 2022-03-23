package app

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/core/permissions"
	"github.com/yitsushi/go-misskey/models"
)

// CreateRequest represents a Create request.
type CreateRequest struct {
	Name        string                   `json:"name"`
	Description string                   `json:"description"`
	Permission  []permissions.Permission `json:"permission"`
	CallbackURL core.String              `json:"callbackUrl,omitempty"`
}

// Validate request.
func (r CreateRequest) Validate() error {
	if r.Name == "" {
		return core.RequestValidationError{
			Request: r,
			Message: core.UndefinedRequiredField,
			Field:   "Name",
		}
	}

	if r.Description == "" {
		return core.RequestValidationError{
			Request: r,
			Message: core.UndefinedRequiredField,
			Field:   "Description",
		}
	}

	if len(r.Permission) == 0 {
		return core.RequestValidationError{
			Request: r,
			Message: core.UndefinedRequiredField,
			Field:   "Permission",
		}
	}

	return nil
}

// Create clip.
func (s *Service) Create(request CreateRequest) (models.App, error) {
	var response models.App
	err := s.Call(
		&core.JSONRequest{Request: &request, Path: "/app/create"},
		&response,
	)

	return response, err
}
