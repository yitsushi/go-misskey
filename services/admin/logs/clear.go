package logs

import "github.com/yitsushi/go-misskey/core"

// ClearRequest represents a Clear request.
type ClearRequest struct{}

// Validate the request.
func (r ClearRequest) Validate() error {
	return nil
}

// Clear logs.
func (s *Service) Clear() error {
	return s.Call(
		&core.JSONRequest{Request: &ClearRequest{}, Path: "/admin/delete-logs"},
		&core.EmptyResponse{},
	)
}
