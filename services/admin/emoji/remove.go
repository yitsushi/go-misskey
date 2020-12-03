package emoji

import (
	"github.com/yitsushi/go-misskey/core"
)

// RemoveRequest represents a Remove request.
type RemoveRequest struct {
	ID string `json:"id"`
}

// Validate the request.
func (r RemoveRequest) Validate() error {
	if r.ID == "" {
		return core.RequestValidationError{
			Request: r,
			Message: core.UndefinedRequiredField,
			Field:   "ID",
		}
	}

	return nil
}

// Remove an emoji.
func (s *Service) Remove(id string) error {
	return s.Call(
		&core.JSONRequest{Request: &RemoveRequest{ID: id}, Path: "/admin/emoji/remove"},
		&core.EmptyResponse{},
	)
}
