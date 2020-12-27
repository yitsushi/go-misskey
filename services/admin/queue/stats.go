package queue

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// StatsRequest represents an Stats request.
type StatsRequest struct{}

// Validate the request.
func (r StatsRequest) Validate() error {
	return nil
}

// Stats shows all queues with stats.
func (s *Service) Stats() (models.QueueStats, error) {
	var response models.QueueStats

	err := s.Call(
		&core.JSONRequest{Request: &StatsRequest{}, Path: "/admin/queue/stats"},
		&response,
	)

	return response, err
}
