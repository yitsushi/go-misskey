package groups

import (
	"github.com/yitsushi/go-misskey/core"
)

// KickRequest represents an Kick request.
type KickRequest struct {
	GroupID string `json:"groupId"`
	UserID  string `json:"userId"`
}

// Validate the request.
func (r KickRequest) Validate() error {
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

// Kick group.
func (s *Service) Kick(groupID, userID string) error {
	request := KickRequest{GroupID: groupID, UserID: userID}
	err := s.Call(
		&core.JSONRequest{Request: &request, Path: "/users/groups/pull"},
		&core.EmptyResponse{},
	)

	return err
}
