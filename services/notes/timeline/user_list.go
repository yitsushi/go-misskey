package timeline

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// UserListRequest represents an UserList request.
type UserListRequest struct {
	ListID                string `json:"listId"`
	Limit                 uint   `json:"limit"`
	SinceID               string `json:"sinceId,omitempty"`
	UntilID               string `json:"untilId,omitempty"`
	SinceDate             uint64 `json:"sinceDate,omitempty"`
	UntilDate             uint64 `json:"untilDate,omitempty"`
	IncludeMyRenotes      bool   `json:"includeMyRenotes"`
	IncludeRenotedMyNotes bool   `json:"includeRenotedMyNotes"`
	IncludeLocalRenotes   bool   `json:"includeLocalRenotes"`
	OnlyWithFiles         bool   `json:"withFiles"`
}

// Validate the request.
func (r UserListRequest) Validate() error {
	if r.ListID == "" {
		return core.RequestValidationError{
			Request: r,
			Message: core.UndefinedRequiredField,
			Field:   "ListID",
		}
	}

	if r.Limit < 1 || r.Limit > maxLimit {
		return core.RequestValidationError{
			Request: r,
			Message: core.NewRangeError(1, maxLimit),
			Field:   "Limit",
		}
	}

	return nil
}

// UserList is the endpoint for /notes/user-list-timeline.
func (s *Service) UserList(request UserListRequest) ([]models.Note, error) {
	var response []models.Note

	err := s.Call(
		&core.JSONRequest{Request: &request, Path: "/notes/user-list-timeline"},
		&response,
	)

	return response, err
}
