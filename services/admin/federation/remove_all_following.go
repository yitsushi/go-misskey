package federation

import "github.com/yitsushi/go-misskey/core"

// RemoveAllFollowingRequest represents a Clear request.
type RemoveAllFollowingRequest struct {
	Host string `json:"host"`
}

// Validate the request.
func (r RemoveAllFollowingRequest) Validate() error {
	if r.Host == "" {
		return core.RequestValidationError{
			Request: r,
			Message: core.UndefinedRequiredField,
			Field:   "Host",
		}
	}

	return nil
}

// RemoveAllFollowing logs.
func (s *Service) RemoveAllFollowing(request RemoveAllFollowingRequest) error {
	return s.Call(
		&core.JSONRequest{Request: &request, Path: "/admin/federation/remove-all-following"},
		&core.EmptyResponse{},
	)
}
