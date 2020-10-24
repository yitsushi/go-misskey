package hashtags

import (
	"github.com/yitsushi/go-misskey/core"
)

// SearchRequest represents an Search request.
type SearchRequest struct {
	Limit  uint   `json:"limit"`
	Query  string `json:"query"`
	Offset int64  `json:"offset"`
}

// SearchOptions are all the options you can play with.
type SearchOptions struct {
	Limit  uint
	Query  string
	Offset int64
}

// Validate options.
func (options *SearchOptions) Validate() error {
	if options.Query == "" {
		return core.MissingOptionsError{
			Endpoint:      "Hashtags/Search",
			Struct:        "SearchOptions",
			MissingFields: []string{"Query"},
		}
	}

	return nil
}

// Search endpoint.
//
// The Query parameter is an SQL LIKE query.
// For example, you are looking for all the hashtags that
// starts with 'hack', the Query should be 'hack%'.
func (s *Service) Search(options *SearchOptions) ([]string, error) {
	var response []string

	if options == nil {
		return response, core.MissingOptionsError{
			Endpoint: "Hashtags/Search",
			Struct:   "SearchOptions",
		}
	}

	err := options.Validate()
	if err != nil {
		return response, err
	}

	request := SearchRequest{
		Limit:  options.Limit,
		Query:  options.Query,
		Offset: options.Offset,
	}

	err = s.Call(
		&core.BaseRequest{Request: &request, Path: "/hashtags/search"},
		&response,
	)

	return response, err
}
