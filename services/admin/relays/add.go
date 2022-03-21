package relays

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// AddRequest represents an Add relay request.
type AddRequest struct {
	Inbox string `json:"inbox"`
}

// Validate the request.
func (r AddRequest) Validate() error {
	if r.Inbox == "" {
		return core.RequestValidationError{
			Request: r,
			Message: core.UndefinedRequiredField,
			Field:   "Inbox",
		}
	}

	return nil
}

// Add a relay.
func (s *Service) Add(inbox string) (models.Relay, error) {
	response := models.Relay{}
	request := AddRequest{Inbox: inbox}
	err := s.Call(
		&core.JSONRequest{Request: &request, Path: "/admin/relays/add"},
		&response,
	)

	return response, err
}
