package federation

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// FollowRequest is the request used by the followers and following commands.
type FollowRequest struct {
	Host    string `json:"host"`
	SinceID string `json:"sinceId"`
	UntilID string `json:"untilId"`
	Limit   int    `json:"limit"`
}

// Follow is a single follower or following record.
type Follow struct {
	ID         string      `json:"id"`
	CreatedAt  string      `json:"createdAt"`
	FolloweeID string      `json:"followeeId"`
	Followee   models.User `json:"followee"`
	FollowerID string      `json:"followerId"`
	Follower   models.User `json:"follower"`
}

// FollowResponse contains a list of followers or followings.
type FollowResponse struct {
	Follows []Follow
}

// Followers lists all followers.
func (s *Service) Followers(request FollowRequest) (FollowResponse, error) {
	var response FollowResponse
	err := s.Call(
		&core.BaseRequest{Request: request, Path: "/federation/followers"},
		&response,
	)

	return response, err
}

// Following lists all followings.
func (s *Service) Following(request FollowRequest) (FollowResponse, error) {
	var response FollowResponse
	err := s.Call(
		&core.BaseRequest{Request: request, Path: "/federation/following"},
		&response,
	)

	return response, err
}
