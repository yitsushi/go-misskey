package polls

import (
	"github.com/yitsushi/go-misskey/core"
)

// RecommendationRequest represents an Recommendation request.
type RecommendationRequest struct {
	Limit  uint   `json:"limit"`
	Offset uint64 `json:"offset"`
}

// Validate the request.
func (r RecommendationRequest) Validate() error {
	return nil
}

// Recommendation endpoint.
func (s *Service) Recommendation(request RecommendationRequest) error {
	return core.NotImplementedYet{
		Reason: "We don't know the response structure",
	}
}
