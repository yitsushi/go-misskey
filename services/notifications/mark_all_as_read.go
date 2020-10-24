package notifications

import (
	"github.com/yitsushi/go-misskey/core"
)

// MarkAllAsReadRequest represents an MarkAllAsRead request.
type MarkAllAsReadRequest struct{}

// MarkAllAsRead endpoint.
func (s *Service) MarkAllAsRead() error {
	err := s.Call(
		&core.BaseRequest{Request: &MarkAllAsReadRequest{}, Path: "/notifications/mark-all-as-read"},
		&core.DummyResponse{},
	)

	return err
}
