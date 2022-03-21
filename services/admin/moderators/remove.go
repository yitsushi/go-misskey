package moderators

import (
	"github.com/yitsushi/go-misskey/core"
)

// RemoveRequest represents an Remove moderator request.
type RemoveRequest struct {
	UserID string `json:"userId"`
}

// Validate the request.
func (r RemoveRequest) Validate() error {
	if r.UserID == "" {
		return core.RequestValidationError{
			Request: r,
			Message: core.UndefinedRequiredField,
			Field:   "UserID",
		}
	}

	return nil
}

// Remove a moderator.
func (s *Service) Remove(userID string) error {
	request := RemoveRequest{UserID: userID}
	err := s.Call(
		&core.JSONRequest{Request: &request, Path: "/admin/moderators/remove"},
		&core.EmptyResponse{},
	)

	return err
}
