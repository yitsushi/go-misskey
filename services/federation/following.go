package federation

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// followingResponse contains a list of followings.
type followingResponse struct {
	Followings []models.Following
}

// followingRequest contains request information for the followings call.
type followingRequest struct {
	Host    string `json:"host"`
	SinceID string `json:"sinceId"`
	UntilID string `json:"untilId"`
	Limit   int    `json:"limit"`
}

// FollowingOptions contains options passed to the following call.
type FollowingOptions struct {
	Host    string
	SinceID string
	UntilID string
	Limit   int
}

// Following lists all followings.
func (s *Service) Following(options FollowingOptions) ([]models.Following, error) {
	var response followingResponse

	request := followingRequest{
		Host:    options.Host,
		SinceID: options.SinceID,
		UntilID: options.UntilID,
		Limit:   options.Limit,
	}

	err := s.Call(
		&core.BaseRequest{Request: request, Path: "/federation/following"},
		&response,
	)

	return response.Followings, err
}
