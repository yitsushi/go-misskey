package notifications

import (
	"github.com/yitsushi/go-misskey/core"
)

// MarkAllAsReadRequest represents an MarkAllAsRead request.
type MarkAllAsReadRequest struct{}

// Validate the request.
func (r MarkAllAsReadRequest) Validate() error {
	return nil
}

// MarkAllAsRead endpoint.
func (s *Service) MarkAllAsRead() error {
	err := s.Call(
		&core.JSONRequest{
			Request: &MarkAllAsReadRequest{},
			Path:    "/notifications/mark-all-as-read",
		},
		&core.DummyResponse{},
	)

	return err
}
