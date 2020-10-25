package federation

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// FollowersRequest contains request information to obtain followers.
type FollowersRequest struct {
	Host    string `json:"host"`
	SinceID string `json:"sinceId"`
	UntilID string `json:"untilId"`
	Limit   int    `json:"limit"`
}

// Validate the request.
func (r *FollowersRequest) Validate() error {
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

// Followers lists all followers.
func (s *Service) Followers(request FollowersRequest) ([]models.FollowStatus, error) {
	var response []models.FollowStatus

	err := s.Call(
		&core.JSONRequest{Request: &request, Path: "/federation/followers"},
		&response,
	)

	return response, err
}
