package federation

import "github.com/yitsushi/go-misskey/core"

// UpdateRemoteUserRequest represents a Clear request.
type UpdateRemoteUserRequest struct {
	UserID string `json:"userId"`
}

// Validate the request.
func (r UpdateRemoteUserRequest) Validate() error {
	if r.UserID == "" {
		return core.RequestValidationError{
			Request: r,
			Message: core.UndefinedRequiredField,
			Field:   "UserID",
		}
	}

	return nil
}

// UpdateRemoteUser logs.
func (s *Service) UpdateRemoteUser(request UpdateRemoteUserRequest) error {
	return s.Call(
		&core.JSONRequest{Request: &request, Path: "/admin/update-remote-user"},
		&core.EmptyResponse{},
	)
}
