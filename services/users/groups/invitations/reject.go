package invitations

import (
	"github.com/yitsushi/go-misskey/core"
)

// RejectRequest represents an Reject request.
type RejectRequest struct {
	InvitationID string `json:"invitationId"`
}

// Validate the request.
func (r RejectRequest) Validate() error {
	if r.InvitationID == "" {
		return core.RequestValidationError{
			Request: r,
			Message: core.UndefinedRequiredField,
			Field:   "InvitationID",
		}
	}

	return nil
}

// Reject group.
func (s *Service) Reject(invitationID string) error {
	request := RejectRequest{InvitationID: invitationID}
	err := s.Call(
		&core.JSONRequest{Request: &request, Path: "/users/groups/invitations/reject"},
		&core.EmptyResponse{},
	)

	return err
}
