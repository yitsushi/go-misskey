package following

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// DeleteRequest is the request structure to delete a following.
type DeleteRequest struct {
	UserID string `json:"userId"`
}

// Validate request.
func (r DeleteRequest) Validate() error {
	if r.UserID == "" {
		return core.RequestValidationError{
			Request: r,
			Message: core.UndefinedRequiredField,
			Field:   "UserID",
		}
	}

	return nil
}

// Delete following endpoint.
func (s *Service) Delete(userID string) (models.User, error) {
	var response models.User

	request := &DeleteRequest{UserID: userID}

	err := s.Call(
		&core.JSONRequest{Request: request, Path: "/following/delete"},
		&response,
	)

	return response, err
}
