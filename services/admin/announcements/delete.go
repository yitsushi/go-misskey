package announcements

import (
	"github.com/yitsushi/go-misskey/core"
)

// DeleteRequest represents an Delete Announcement request.
type DeleteRequest struct {
	ID string `json:"id"`
}

// Validate the request.
func (r DeleteRequest) Validate() error {
	if r.ID == "" {
		return core.RequestValidationError{
			Request: r,
			Message: core.UndefinedRequiredField,
			Field:   "ID",
		}
	}

	return nil
}

// Delete an announcement.
func (s *Service) Delete(id string) error {
	err := s.Call(
		&core.JSONRequest{
			Request: &DeleteRequest{
				ID: id,
			},
			Path: "/admin/announcements/delete",
		},
		&core.EmptyResponse{},
	)

	return err
}
