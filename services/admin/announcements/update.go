package announcements

import (
	"github.com/yitsushi/go-misskey/core"
)

// UpdateRequest represents an Update Announcement request.
type UpdateRequest struct {
	ID       string      `json:"id"`
	Title    string      `json:"title"`
	Text     string      `json:"text"`
	ImageURL core.String `json:"imageUrl,omitempty"`
}

// Validate the request.
func (r UpdateRequest) Validate() error {
	if r.ID == "" {
		return core.RequestValidationError{
			Request: r,
			Message: core.UndefinedRequiredField,
			Field:   "ID",
		}
	}

	if r.Title == "" {
		return core.RequestValidationError{
			Request: r,
			Message: core.UndefinedRequiredField,
			Field:   "Title",
		}
	}

	if r.Text == "" {
		return core.RequestValidationError{
			Request: r,
			Message: core.UndefinedRequiredField,
			Field:   "Text",
		}
	}

	if r.ImageURL != nil && core.StringValue(r.ImageURL) == "" {
		return core.RequestValidationError{
			Request: r,
			Message: core.UndefinedRequiredField,
			Field:   "ImageURL",
		}
	}

	return nil
}

// Update an announcement.
func (s *Service) Update(request UpdateRequest) error {
	err := s.Call(
		&core.JSONRequest{Request: &request, Path: "/admin/announcements/update"},
		&core.EmptyResponse{},
	)

	return err
}
