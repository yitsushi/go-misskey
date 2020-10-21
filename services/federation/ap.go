package federation

import "github.com/yitsushi/go-misskey/core"

// ShowRequest is the request used by the show command.
// Contains a single URI.
type ShowRequest struct {
	URI string `json:"uri"`
}

// ShowResponse is an empty response for the Show command.
type ShowResponse struct {
}

// Show shows the ActivityPub object specified by the URI.
func (s *Service) Show(request ShowRequest) error {
	var response ShowResponse
	err := s.Call(
		&core.BaseRequest{Request: request, Path: "/ap/show"},
		&response,
	)

	return err
}
