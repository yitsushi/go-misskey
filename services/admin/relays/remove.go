package relays

import (
	"github.com/yitsushi/go-misskey/core"
)

// RemoveRequest represents a Remove relay request.
type RemoveRequest struct {
	Inbox string `json:"inbox"`
}

// Validate the request.
func (r RemoveRequest) Validate() error {
	if r.Inbox == "" {
		return core.RequestValidationError{
			Request: r,
			Message: core.UndefinedRequiredField,
			Field:   "Inbox",
		}
	}

	return nil
}

// Remove a moderator.
func (s *Service) Remove(inbox string) error {
	request := RemoveRequest{Inbox: inbox}
	err := s.Call(
		&core.JSONRequest{Request: &request, Path: "/admin/relays/remove"},
		&core.EmptyResponse{},
	)

	return err
}
