package emoji

import (
	"github.com/yitsushi/go-misskey/core"
)

// CopyRequest represents a Copy request.
type CopyRequest struct {
	EmojiID string `json:"emojiId"`
}

// Validate the request.
func (r CopyRequest) Validate() error {
	if r.EmojiID == "" {
		return core.RequestValidationError{
			Request: r,
			Message: core.UndefinedRequiredField,
			Field:   "EmojiID",
		}
	}

	return nil
}

type copyResponse struct {
	ID string `json:"id"`
}

// Copy an emoji.
func (s *Service) Copy(request CopyRequest) (string, error) {
	var response copyResponse

	err := s.Call(
		&core.JSONRequest{Request: &request, Path: "/admin/emoji/copy"},
		&response,
	)

	return response.ID, err
}
