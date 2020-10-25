package clips

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// ShowRequest represents an Show request.
type ShowRequest struct {
	ClipID string `json:"clipId"`
}

// Validate the request.
func (r ShowRequest) Validate() error {
	if r.ClipID == "" {
		return core.RequestValidationError{
			Request: r,
			Message: core.UndefinedRequiredField,
			Field:   "ClipID",
		}
	}

	return nil
}

// Show clips.
func (s *Service) Show(request ShowRequest) (models.Clip, error) {
	var response models.Clip
	err := s.Call(
		&core.JSONRequest{Request: &request, Path: "/clips/show"},
		&response,
	)

	return response, err
}
