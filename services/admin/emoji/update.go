package emoji

import (
	"github.com/yitsushi/go-misskey/core"
)

// UpdateRequest represents an Update request.
type UpdateRequest struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Category string   `json:"category,omitempty"`
	Aliases  []string `json:"aliases"`
}

// Validate the request.
func (r UpdateRequest) Validate() error {
	if r.ID == "" {
		return core.RequestValidationError{
			Request: r,
			Message: core.UndefinedRequiredField,
			Field:   "ID",
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

// Update an emoji.
func (s *Service) Update(request UpdateRequest) error {
	if request.Aliases == nil {
		request.Aliases = []string{}
	}

	return s.Call(
		&core.JSONRequest{Request: &request, Path: "/admin/emoji/update"},
		&core.EmptyResponse{},
	)
}
