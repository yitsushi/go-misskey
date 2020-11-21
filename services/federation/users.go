package federation

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// UsersRequest contains request information to obtain users.
type UsersRequest struct {
	Host    string `json:"host"`
	SinceID string `json:"sinceId"`
	UntilID string `json:"untilId"`
	Limit   uint   `json:"limit"`
}

// Validate the request.
func (r UsersRequest) Validate() error {
	if r.Host == "" {
		return core.RequestValidationError{
			Request: r,
			Message: core.UndefinedRequiredField,
			Field:   "Host",
		}
	}

	if r.Limit < 1 || r.Limit > 100 {
		return core.RequestValidationError{
			Request: r,
			Message: core.NewRangeError(1, 100),
			Field:   "Limit",
		}
	}

	return nil
}

// Users will list all users for a federation with a given host.
func (s *Service) Users(request UsersRequest) ([]models.User, error) {
	var response []models.User

	err := s.Call(
		&core.JSONRequest{Request: &request, Path: "/federation/users"},
		&response,
	)

	return response, err
}
