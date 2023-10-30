package instance

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// ServerInfoRequest represents a Clear request.
type ServerInfoRequest struct{}

// Validate the request.
func (r ServerInfoRequest) Validate() error {
	return nil
}

// ServerInfo gets the server info.
func (s *Service) ServerInfo() (models.ServerInfo, error) {
	var response models.ServerInfo
	err := s.Call(
		&core.JSONRequest{Request: &ServerInfoRequest{}, Path: "/admin/server-info"},
		&response,
	)

	return response, err
}
