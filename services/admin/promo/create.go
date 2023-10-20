package promo

import (
	"github.com/yitsushi/go-misskey/core"
)

// CreateRequest is the request object for a Create request.
type CreateRequest struct {
	NoteID    string `json:"noteId"`    // required: <misskey:id>
	ExpiresAt int64  `json:"expiresAt"` // required: time.Now().Add(0,86400000,0).Unix()
}

// Validate the request.
func (r CreateRequest) Validate() error {
	if r.NoteID == "" {
		return core.RequestValidationError{
			Request: r,
			Message: core.UndefinedRequiredField,
			Field:   "noteId",
		}
	}

	return nil
}

// Create is the endpoint to create a promo.
func (s *Service) Create(request CreateRequest) error {
	err := s.Call(
		&core.JSONRequest{
			Request: &request,
			Path:    "/admin/promo/create",
		},
		&core.EmptyResponse{},
	)

	return err
}
