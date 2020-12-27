package federation

import "github.com/yitsushi/go-misskey/core"

// DeleteAllFilesRequest represents a Clear request.
type DeleteAllFilesRequest struct {
	Host string `json:"host"`
}

// Validate the request.
func (r DeleteAllFilesRequest) Validate() error {
	if r.Host == "" {
		return core.RequestValidationError{
			Request: r,
			Message: core.UndefinedRequiredField,
			Field:   "Host",
		}
	}

	return nil
}

// DeleteAllFiles logs.
func (s *Service) DeleteAllFiles(request DeleteAllFilesRequest) error {
	return s.Call(
		&core.JSONRequest{Request: &request, Path: "/admin/federation/delete-all-files"},
		&core.EmptyResponse{},
	)
}
