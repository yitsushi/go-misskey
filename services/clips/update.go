package clips

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// UpdateRequest represents a Update request.
type UpdateRequest struct {
	ClipID string `json:"clipId"`
	Name   string `json:"name"`
}

// Validate request.
func (r UpdateRequest) Validate() error {
	if len(r.Name) > MaximumNameLength {
		return core.RequestValidationError{
			Request: r,
			Message: core.ExceedMaximumLengthError,
			Field:   "Name",
		}
	}

	if r.ClipID == "" {
		return core.RequestValidationError{
			Request: r,
			Message: core.UndefinedRequiredField,
			Field:   "ClipID",
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

// Update clip.
func (s *Service) Update(request UpdateRequest) (models.Clip, error) {
	var response models.Clip
	err := s.Call(
		&core.JSONRequest{Request: &request, Path: "/clips/update"},
		&response,
	)

	return response, err
}
