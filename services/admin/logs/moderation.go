package logs

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// ModerationRequest represents an admin/show-moderation-logs request.
type ModerationRequest struct {
	Limit   uint   `json:"limit,omitempty"`
	SinceID string `json:"sinceId,omitempty"`
	UntilID string `json:"untilId,omitempty"`
}

// Validate the request.
func (r ModerationRequest) Validate() error {
	return nil
}

// Moderation logs.
func (s *Service) Moderation() ([]models.ModerationLog, error) {
	var response []models.ModerationLog

	err := s.Call(
		&core.JSONRequest{Request: &ModerationRequest{}, Path: "/admin/show-moderation-logs"},
		&response,
	)

	return response, err
}
