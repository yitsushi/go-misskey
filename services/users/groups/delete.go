package groups

import (
	"github.com/yitsushi/go-misskey/core"
)

// DeleteRequest represents an Delete request.
type DeleteRequest struct {
	GroupID string `json:"groupId"`
}

// Validate the request.
func (r DeleteRequest) Validate() error {
	if r.GroupID == "" {
		return core.RequestValidationError{
			Request: r,
			Message: core.UndefinedRequiredField,
			Field:   "GroupID",
		}
	}

	return nil
}

// Delete group.
func (s *Service) Delete(groupID string) error {
	request := DeleteRequest{GroupID: groupID}
	err := s.Call(
		&core.JSONRequest{Request: &request, Path: "/users/groups/delete"},
		&core.EmptyResponse{},
	)

	return err
}
