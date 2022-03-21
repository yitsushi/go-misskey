package groups

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// UpdateRequest represents an Update request.
type UpdateRequest struct {
	GroupID string `json:"groupId"`
	Name    string `json:"name"`
}

// Validate the request.
func (r UpdateRequest) Validate() error {
	if r.GroupID == "" {
		return core.RequestValidationError{
			Request: r,
			Message: core.UndefinedRequiredField,
			Field:   "GroupID",
		}
	}

	if r.Name == "" {
		return core.RequestValidationError{
			Request: r,
			Message: core.UndefinedRequiredField,
			Field:   "Name",
		}
	}

	return nil
}

// Update group.
func (s *Service) Update(groupID, name string) (models.Group, error) {
	request := UpdateRequest{GroupID: groupID, Name: name}
	response := models.Group{}
	err := s.Call(
		&core.JSONRequest{Request: &request, Path: "/users/groups/update"},
		&response,
	)

	return response, err
}
