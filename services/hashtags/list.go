package hashtags

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// ListRequest represents an List request.
type ListRequest struct {
	Limit                    uint   `json:"limit"`
	AttachedToUserOnly       bool   `json:"attachedToUserOnly"`
	AttachedToLocalUserOnly  bool   `json:"attachedToLocalUserOnly"`
	AttachedToRemoteUserOnly bool   `json:"attachedToRemoteUserOnly"`
	Sort                     string `json:"sort"`
}

// Validate the request.
func (r ListRequest) Validate() error {
	if r.Sort == "" {
		return core.RequestValidationError{
			Request: r,
			Message: core.UndefinedRequiredField,
			Field:   "Sort",
		}
	}

	return nil
}

// List endpoint.
func (s *Service) List(request ListRequest) ([]models.Hashtag, error) {
	var response []models.Hashtag
	err := s.Call(
		&core.JSONRequest{Request: &request, Path: "/hashtags/list"},
		&response,
	)

	return response, err
}
