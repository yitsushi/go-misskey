package admin

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/services/admin/announcements"
	"github.com/yitsushi/go-misskey/services/admin/emoji"
	"github.com/yitsushi/go-misskey/services/admin/federation"
	"github.com/yitsushi/go-misskey/services/admin/logs"
	"github.com/yitsushi/go-misskey/services/admin/moderation"
	"github.com/yitsushi/go-misskey/services/admin/queue"
	"github.com/yitsushi/go-misskey/services/admin/users"
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

// Users contains all endpoints for users.
func (s *Service) Users() *users.Service {
	return users.NewService(s.Call)
}

// Emoji contains all endpoints for emoji.
func (s *Service) Emoji() *emoji.Service {
	return emoji.NewService(s.Call)
}

// Queue contains all endpoints for queue.
func (s *Service) Queue() *queue.Service {
	return queue.NewService(s.Call)
}

// Logs contains all endpoints for logs.
func (s *Service) Logs() *logs.Service {
	return logs.NewService(s.Call)
}

// Federation contains all endpoints for federation.
func (s *Service) Federation() *federation.Service {
	return federation.NewService(s.Call)
}
