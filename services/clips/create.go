package clips

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// CreateRequest represents a Create request.
type CreateRequest struct {
	Name string `json:"name"`
}

// Validate request.
func (r CreateRequest) Validate() error {
	if len(r.Name) > maximumNameLength {
		return core.RequestValidationError{
			Request: r,
			Message: core.ExceedMaximumLengthError,
			Field:   "Name",
		}
	}

	if r.Name == "" {
		return core.RequestValidationError{
			Request: r,
			Message: core.UndefinedRequiredField,
			Field:   "Name",
		}
	}

	return nil
}

// Create clip.
func (s *Service) Create(request CreateRequest) (models.Clip, error) {
	var response models.Clip
	err := s.Call(
		&core.JSONRequest{Request: &request, Path: "/clips/create"},
		&response,
	)

	return response, err
}
