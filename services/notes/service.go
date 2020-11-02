package notes

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/services/notes/reactions"
	"github.com/yitsushi/go-misskey/services/notes/timeline"
)

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
