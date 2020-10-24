package federation

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// followersResponse contains a list of followers.
type followersResponse struct {
	Followers []models.Followers
}

// followersRequest contains request information for the followers call.
type followersRequest struct {
	Host    string `json:"host"`
	SinceID string `json:"sinceId"`
	UntilID string `json:"untilId"`
	Limit   int    `json:"limit"`
}

// FollowersOptions contains options passed to the followers call.
type FollowersOptions struct {
	Host    string
	SinceID string
	UntilID string
	Limit   int
}

// Followers lists all followers.
func (s *Service) Followers(options FollowersOptions) ([]models.Followers, error) {
	var response followersResponse

	request := followersRequest{
		Host:    options.Host,
		SinceID: options.SinceID,
		UntilID: options.UntilID,
		Limit:   options.Limit,
	}

	err := s.Call(
		&core.BaseRequest{Request: request, Path: "/federation/followers"},
		&response,
	)

	return response.Followers, err
}
