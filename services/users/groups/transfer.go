package groups

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// TransferRequest represents an Transfer request.
type TransferRequest struct {
	GroupID string `json:"groupId"`
	UserID  string `json:"userId"`
}

// Validate the request.
func (r TransferRequest) Validate() error {
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

// Transfer group.
func (s *Service) Transfer(groupID, userID string) (models.Group, error) {
	request := TransferRequest{GroupID: groupID, UserID: userID}
	response := models.Group{}
	err := s.Call(
		&core.JSONRequest{Request: &request, Path: "/users/groups/transfer"},
		&response,
	)

	return response, err
}
