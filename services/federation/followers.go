package federation

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// followersResponse contains a list of followers.
type followersResponse struct {
	Followers []models.Followers
}

// FollowersRequest contains request information for the followers call.
type FollowersRequest struct {
	Host    string `json:"host"`
	SinceID string `json:"sinceId"`
	UntilID string `json:"untilId"`
	Limit   int    `json:"limit"`
}

// Validate the request.
func (r *FollowersRequest) Validate() error {
	return nil
}

// Followers lists all followers.
func (s *Service) Followers(request *FollowersRequest) ([]models.Followers, error) {
	var response followersResponse

	err := s.Call(
		&core.JSONRequest{Request: request, Path: "/federation/followers"},
		&response,
	)

	return response.Followers, err
}
