package hashtags

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

const maxLimit = 100

// UserState of user.
type UserState string

const (
	// StateAll for all users.
	StateAll UserState = "all"
	// StateAlive for only alive users.
	StateAlive UserState = "alive"

	// DefaultState is the default state.
	DefaultState UserState = StateAll
	// DefaultOrigin is the default origin.
	DefaultOrigin models.UserOrigin = models.OriginLocal
)

// UsersRequest represents an Users request.
type UsersRequest struct {
	Tag    string            `json:"tag"`
	Limit  uint              `json:"limit"`
	Sort   string            `json:"sort"`
	State  UserState         `json:"state"`
	Origin models.UserOrigin `json:"origin"`
}

// Validate options.
func (r UsersRequest) Validate() error {
	if r.Sort == "" {
		return core.RequestValidationError{
			Request: r,
			Message: core.UndefinedRequiredField,
			Field:   "Sort",
		}
	}

	if r.Tag == "" {
		return core.RequestValidationError{
			Request: r,
			Message: core.UndefinedRequiredField,
			Field:   "Tag",
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

// Users endpoint.
func (s *Service) Users(request UsersRequest) ([]models.User, error) {
	var response []models.User
	err := s.Call(
		&core.JSONRequest{Request: &request, Path: "/hashtags/users"},
		&response,
	)

	return response, err
}
