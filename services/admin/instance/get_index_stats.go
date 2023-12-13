package instance

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// GetIndexStatsRequest represents a GetIndexStats request.
type GetIndexStatsRequest struct{}

// Validate the request.
func (r GetIndexStatsRequest) Validate() error {
	return nil
}

// GetIndexStats gets the index statistics.
func (s *Service) GetIndexStats() (models.IndexStats, error) {
	var response models.IndexStats
	err := s.Call(
		&core.JSONRequest{Request: &GetIndexStatsRequest{}, Path: "/admin/get-index-stats"},
		&response,
	)

	return response, err
}
