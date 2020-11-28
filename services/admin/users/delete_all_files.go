package users

import (
	"github.com/yitsushi/go-misskey/core"
)

// DeleteAllFilesRequest is the request object for a DeleteAllFiles request.
type DeleteAllFilesRequest struct {
	UserID string `json:"userId"`
}

// Validate the request.
func (r DeleteAllFilesRequest) Validate() error {
	if r.UserID == "" {
		return core.RequestValidationError{
			Request: r,
			Message: core.UndefinedRequiredField,
			Field:   "UserID",
		}
	}

	return nil
}

// DeleteAllFiles is the endpoint to get a User.
func (s *Service) DeleteAllFiles(userID string) error {
	return s.Call(
		&core.JSONRequest{
			Request: &DeleteAllFilesRequest{
				UserID: userID,
			},
			Path: "/admin/delete-all-files-of-a-user",
		},
		&core.EmptyResponse{},
	)
}
