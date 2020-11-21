package announcements

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// CreateRequest represents an Create Announcement request.
type CreateRequest struct {
	Title    string      `json:"title"`
	Text     string      `json:"text"`
	ImageURL core.String `json:"imageUrl,omitempty"`
}

// Validate the request.
func (r CreateRequest) Validate() error {
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

// Create an announcement.
func (s *Service) Create(request CreateRequest) (models.Announcement, error) {
	var response models.Announcement
	err := s.Call(
		&core.JSONRequest{Request: &request, Path: "/admin/announcements/create"},
		&response,
	)

	return response, err
}
