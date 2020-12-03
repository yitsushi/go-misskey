package emoji

import (
	"github.com/yitsushi/go-misskey/core"
)

// AddRequest represents an Add request.
type AddRequest struct {
	FileID string `json:"fileId"`
}

// Validate the request.
func (r AddRequest) Validate() error {
	if r.FileID == "" {
		return core.RequestValidationError{
			Request: r,
			Message: core.UndefinedRequiredField,
			Field:   "FileID",
		}
	}

	return nil
}

type addResponse struct {
	ID string `json:"id"`
}

// Add an emoji.
func (s *Service) Add(request AddRequest) (string, error) {
	var response addResponse

	err := s.Call(
		&core.JSONRequest{Request: &request, Path: "/admin/emoji/add"},
		&response,
	)

	return response.ID, err
}
