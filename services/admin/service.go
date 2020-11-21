package admin

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/services/admin/announcements"
	"github.com/yitsushi/go-misskey/services/admin/moderation"
)

// Service is the base for all the endpoints on this service.
type Service struct {
	Call core.RequestHandlerFunc
}

// NewService creates a new Service instance.
func NewService(requestHandler core.RequestHandlerFunc) *Service {
	return &Service{Call: requestHandler}
}

// Accouncements contains all endpoints under /admin/announcements.
func (s *Service) Accouncements() *announcements.Service {
	return announcements.NewService(s.Call)
}

// Moderation contains all endpoints for moderation.
func (s *Service) Moderation() *moderation.Service {
	return moderation.NewService(s.Call)
}
