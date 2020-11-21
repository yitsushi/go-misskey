package hashtags

import (
	"github.com/yitsushi/go-misskey/core"
)

// SearchRequest represents an Search request.
type SearchRequest struct {
	Limit  uint   `json:"limit"`
	Query  string `json:"query"`
	Offset uint64 `json:"offset"`
}

// Validate options.
func (r SearchRequest) Validate() error {
	if r.Query == "" {
		return core.RequestValidationError{
			Request: r,
			Message: core.UndefinedRequiredField,
			Field:   "Query",
		}
	}

	return nil
}

// Search endpoint.
//
// The Query parameter is an SQL LIKE query.
// For example, you are looking for all the hashtags that
// starts with 'hack', the Query should be 'hack%'.
func (s *Service) Search(request SearchRequest) ([]string, error) {
	var response []string
	err := s.Call(
		&core.JSONRequest{Request: &request, Path: "/hashtags/search"},
		&response,
	)

	return response, err
}
