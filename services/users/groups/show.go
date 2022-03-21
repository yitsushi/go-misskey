package groups

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// ShowRequest represents an Show request.
type ShowRequest struct {
	GroupID string `json:"groupId"`
}

// Validate the request.
func (r ShowRequest) Validate() error {
	if r.GroupID == "" {
		return core.RequestValidationError{
			Request: r,
			Message: core.UndefinedRequiredField,
			Field:   "GroupID",
		}
	}

	return nil
}

// Show group.
func (s *Service) Show(groupID string) (models.Group, error) {
	request := ShowRequest{GroupID: groupID}
	response := models.Group{}
	err := s.Call(
		&core.JSONRequest{Request: &request, Path: "/users/groups/show"},
		&response,
	)

	return response, err
}
