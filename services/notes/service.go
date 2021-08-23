package notes

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/services/notes/favorites"
	"github.com/yitsushi/go-misskey/services/notes/polls"
	"github.com/yitsushi/go-misskey/services/notes/reactions"
	"github.com/yitsushi/go-misskey/services/notes/timeline"
	"github.com/yitsushi/go-misskey/services/notes/watching"
)

const maxLimit = 100

// Service is the base for all the endpoints on this service.
type Service struct {
	Call core.RequestHandlerFunc
}

// NewService creates a new Service instance.
func NewService(requestHandler core.RequestHandlerFunc) *Service {
	return &Service{Call: requestHandler}
}

// Reactions contains all endpoints under /notes/reactions.
func (s *Service) Reactions() *reactions.Service {
	return reactions.NewService(s.Call)
}

// Timeline contains all endpoints related to /notes/timeline.
func (s *Service) Timeline() *timeline.Service {
	return timeline.NewService(s.Call)
}

// Polls contains all endpoints related to /notes/polls.
func (s *Service) Polls() *polls.Service {
	return polls.NewService(s.Call)
}

// Favorites contains all endpoints related to /notes/favorites.
func (s *Service) Favorites() *favorites.Service {
	return favorites.NewService(s.Call)
}

// Watching contains all endpoints related to /notes/watching.
func (s *Service) Watching() *watching.Service {
	return watching.NewService(s.Call)
}
