package federation

import "github.com/yitsushi/go-misskey/core"

// UpdateInstanceRequest represents a Clear request.
type UpdateInstanceRequest struct {
	Host        string `json:"host"`
	IsSuspended bool   `json:"isSuspended"`
}

// Validate the request.
func (r UpdateInstanceRequest) Validate() error {
	if r.Host == "" {
		return core.RequestValidationError{
			Request: r,
			Message: core.UndefinedRequiredField,
			Field:   "Host",
		}
	}

	return nil
}

// UpdateInstance logs.
func (s *Service) UpdateInstance(request UpdateInstanceRequest) error {
	return s.Call(
		&core.JSONRequest{Request: &request, Path: "/admin/federation/update-instance"},
		&core.EmptyResponse{},
	)
}
