package queue

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// DelayedRequest represents an Delayed request.
type DelayedRequest struct{}

// Validate the request.
func (r DelayedRequest) Validate() error {
	return nil
}

// DeliverDelayed lists delayed deliver hosts with count.
func (s *Service) DeliverDelayed() ([]models.Delayed, error) {
	var response [][]interface{}

	err := s.Call(
		&core.JSONRequest{Request: &DelayedRequest{}, Path: "/admin/queue/deliver-delayed"},
		&response,
	)

	delayed := []models.Delayed{}

	for _, item := range response {
		delayed = append(delayed, models.Delayed{
			Host:  item[0].(string),
			Count: int64(item[1].(float64)),
		})
	}

	return delayed, err
}

// InboxDelayed lists delayed inbox hosts with count.
func (s *Service) InboxDelayed() ([]models.Delayed, error) {
	var response [][]interface{}

	err := s.Call(
		&core.JSONRequest{Request: &DelayedRequest{}, Path: "/admin/queue/inbox-delayed"},
		&response,
	)

	delayed := []models.Delayed{}

	for _, item := range response {
		delayed = append(delayed, models.Delayed{
			Host:  item[0].(string),
			Count: int64(item[1].(float64)),
		})
	}

	return delayed, err
}
