package notifications

import (
	"github.com/yitsushi/go-misskey/core"
)

// CreateRequest represents an Create request.
type CreateRequest struct {
	Body   string      `json:"body"`
	Header core.String `json:"header"`
	Icon   core.String `json:"icon"`
}

// CreateOptions are all the options you can play with.
type CreateOptions struct {
	Body   string
	Header core.String
	Icon   core.String
}

// Create endpoint.
func (s *Service) Create(options *CreateOptions) error {
	var response core.DummyResponse

	if options == nil {
		return core.MissingOptionsError{
			Endpoint: "Notifications/Create",
			Struct:   "CreateOptions",
		}
	}

	request := CreateRequest{
		Body:   options.Body,
		Header: options.Header,
		Icon:   options.Icon,
	}

	err := s.Call(
		&core.BaseRequest{Request: &request, Path: "/notifications/create"},
		&response,
	)

	return err
}
