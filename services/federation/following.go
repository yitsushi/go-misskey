package federation

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// followingResponse contains a list of followings.
type followingResponse struct {
	Followings []models.Following
}

// FollowingRequest contains request information for the followings call.
type FollowingRequest struct {
	Host    string `json:"host"`
	SinceID string `json:"sinceId"`
	UntilID string `json:"untilId"`
	Limit   int    `json:"limit"`
}

// Validate the request.
func (r *FollowingRequest) Validate() error {
	return nil
}

// Following lists all followings.
func (s *Service) Following(request *FollowingRequest) ([]models.Following, error) {
	var response followingResponse

	err := s.Call(
		&core.JSONRequest{Request: request, Path: "/federation/following"},
		&response,
	)

	return response.Followings, err
}
