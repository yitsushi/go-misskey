package invitations

import (
	"github.com/yitsushi/go-misskey/core"
)

// AcceptRequest represents an Accept request.
type AcceptRequest struct {
	InvitationID string `json:"invitationId"`
}

// Validate the request.
func (r AcceptRequest) Validate() error {
	if r.InvitationID == "" {
		return core.RequestValidationError{
			Request: r,
			Message: core.UndefinedRequiredField,
			Field:   "InvitationID",
		}
	}

	return nil
}

// Accept group.
func (s *Service) Accept(invitationID string) error {
	request := AcceptRequest{InvitationID: invitationID}
	err := s.Call(
		&core.JSONRequest{Request: &request, Path: "/users/groups/invitations/accept"},
		&core.EmptyResponse{},
	)

	return err
}
