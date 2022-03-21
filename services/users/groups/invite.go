package groups

import (
	"github.com/yitsushi/go-misskey/core"
)

// InviteRequest represents an Invite request.
type InviteRequest struct {
	GroupID string `json:"groupId"`
	UserID  string `json:"userId"`
}

// Validate the request.
func (r InviteRequest) Validate() error {
	if r.GroupID == "" {
		return core.RequestValidationError{
			Request: r,
			Message: core.UndefinedRequiredField,
			Field:   "GroupID",
		}
	}

	if r.UserID == "" {
		return core.RequestValidationError{
			Request: r,
			Message: core.UndefinedRequiredField,
			Field:   "UserID",
		}
	}

	return nil
}

// Invite group.
func (s *Service) Invite(groupID, userID string) error {
	request := InviteRequest{GroupID: groupID, UserID: userID}
	err := s.Call(
		&core.JSONRequest{Request: &request, Path: "/users/groups/invite"},
		&core.EmptyResponse{},
	)

	return err
}
