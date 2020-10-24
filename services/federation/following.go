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

// Following lists all followings.
func (s *Service) Following(host, sinceID, untilID string, limit int) ([]models.Following, error) {
	var response followingResponse

	request := followingRequest{
		Host:    host,
		SinceID: sinceID,
		UntilID: untilID,
		Limit:   limit,
	}

	err := s.Call(
		&core.BaseRequest{Request: request, Path: "/federation/following"},
		&response,
	)

	return response.Followings, err
}
