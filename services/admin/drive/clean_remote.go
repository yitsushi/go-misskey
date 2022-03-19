package drive

import (
	"github.com/yitsushi/go-misskey/core"
)

// CleanRemoteRequest represents a CleanRemote request.
type CleanRemoteRequest struct{}

// Validate the request.
func (r CleanRemoteRequest) Validate() error {
	return nil
}

// CleanRemote drive.
func (s *Service) CleanRemote() error {
	err := s.Call(
		&core.JSONRequest{Request: &CleanRemoteRequest{}, Path: "/admin/drive/clean-remote-files"},
		&core.EmptyResponse{},
	)

	return err
}
