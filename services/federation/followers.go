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

// Followers lists all followers.
func (s *Service) Followers(host, sinceID, untilID string, limit int) ([]models.Followers, error) {
	var response followersResponse

	request := FollowersRequest{
		Host:    host,
		SinceID: sinceID,
		UntilID: untilID,
		Limit:   limit,
	}

	err := s.Call(
		&core.BaseRequest{Request: request, Path: "/federation/followers"},
		&response,
	)

	return response.Followers, err
}
