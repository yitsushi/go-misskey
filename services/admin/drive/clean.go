package drive

import (
	"github.com/yitsushi/go-misskey/core"
)

// CleanRequest represents a Clean request.
type CleanRequest struct{}

// Validate the request.
func (r CleanRequest) Validate() error {
	return nil
}

// Clean local drive.
func (s *Service) Clean() error {
	err := s.Call(
		&core.JSONRequest{Request: &CleanRequest{}, Path: "/admin/drive/cleanup"},
		&core.EmptyResponse{},
	)

	return err
}
