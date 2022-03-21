package groups

import (
	"github.com/yitsushi/go-misskey/core"
)

// LeaveRequest represents an Leave request.
type LeaveRequest struct {
	GroupID string `json:"groupId"`
}

// Validate the request.
func (r LeaveRequest) Validate() error {
	if r.GroupID == "" {
		return core.RequestValidationError{
			Request: r,
			Message: core.UndefinedRequiredField,
			Field:   "GroupID",
		}
	}

	return nil
}

// Leave group.
func (s *Service) Leave(groupID string) error {
	request := LeaveRequest{GroupID: groupID}
	err := s.Call(
		&core.JSONRequest{Request: &request, Path: "/users/groups/leave"},
		&core.EmptyResponse{},
	)

	return err
}
