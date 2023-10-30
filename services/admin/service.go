package admin

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/services/admin/announcements"
	"github.com/yitsushi/go-misskey/services/admin/drive"
	"github.com/yitsushi/go-misskey/services/admin/emoji"
	"github.com/yitsushi/go-misskey/services/admin/federation"
	"github.com/yitsushi/go-misskey/services/admin/instance"
	"github.com/yitsushi/go-misskey/services/admin/logs"
	"github.com/yitsushi/go-misskey/services/admin/moderation"
	"github.com/yitsushi/go-misskey/services/admin/moderators"
	"github.com/yitsushi/go-misskey/services/admin/promo"
	"github.com/yitsushi/go-misskey/services/admin/queue"
	"github.com/yitsushi/go-misskey/services/admin/relays"
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

// Announcements contains all endpoints under /admin/announcements.
func (s *Service) Announcements() *announcements.Service {
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

// Drive contains all endpoints for drive.
func (s *Service) Drive() *drive.Service {
	return drive.NewService(s.Call)
}

// Moderators contains all endpoints for moderators.
func (s *Service) Moderators() *moderators.Service {
	return moderators.NewService(s.Call)
}

// Relays contains all endpoints for relays.
func (s *Service) Relays() *relays.Service {
	return relays.NewService(s.Call)
}

// Promo contains all endpoints for promos.
func (s *Service) Promo() *promo.Service {
	return promo.NewService(s.Call)
}

// Instance contains all endpoints for info about the instance.
func (s *Service) Instance() *instance.Service {
	return instance.NewService(s.Call)
}
