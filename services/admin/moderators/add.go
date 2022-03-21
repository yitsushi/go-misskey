package moderators

import (
	"github.com/yitsushi/go-misskey/core"
)

// AddRequest represents an Add moderator request.
type AddRequest struct {
	UserID string `json:"userId"`
}

// Validate the request.
func (r AddRequest) Validate() error {
	if r.UserID == "" {
		return core.RequestValidationError{
			Request: r,
			Message: core.UndefinedRequiredField,
			Field:   "UserID",
		}
	}

	return nil
}

// Add a moderator.
func (s *Service) Add(userID string) error {
	request := AddRequest{UserID: userID}
	err := s.Call(
		&core.JSONRequest{Request: &request, Path: "/admin/moderators/add"},
		&core.EmptyResponse{},
	)

	return err
}
