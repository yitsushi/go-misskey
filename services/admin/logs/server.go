package logs

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// ServerRequest represents an admin/logs request.
type ServerRequest struct {
	Limit  uint   `json:"limit,omitempty"`
	Level  string `json:"level,omitempty"`
	Domain string `json:"domain,omitempty"`
}

// Validate the request.
func (r ServerRequest) Validate() error {
	return nil
}

// Server logs.
func (s *Service) Server() ([]models.Log, error) {
	var response []models.Log

	err := s.Call(
		&core.JSONRequest{Request: &ServerRequest{}, Path: "/admin/logs"},
		&response,
	)

	return response, err
}
