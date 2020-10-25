package federation

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// FollowingRequest contains request information to obtain followees.
type FollowingRequest struct {
	Host    string `json:"host"`
	SinceID string `json:"sinceId"`
	UntilID string `json:"untilId"`
	Limit   int    `json:"limit"`
}

// Validate the request.
func (r *FollowingRequest) Validate() error {
	if r.Host == "" {
		return core.RequestValidationError{
			Request: r,
			Message: core.UndefinedRequiredField,
			Field:   "Host",
		}
	}
	return nil
}

// Following lists all followings.
func (s *Service) Following(request FollowingRequest) ([]models.FollowStatus, error) {
	var response []models.FollowStatus

	err := s.Call(
		&core.JSONRequest{Request: &request, Path: "/federation/following"},
		&response,
	)

	return response, err
}
