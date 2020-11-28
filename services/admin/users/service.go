package users

import (
	"github.com/yitsushi/go-misskey/core"
)

// UserState is the state of a report.
type UserState string

const (
	// AllState for all users.
	AllState UserState = ""
	// AvailableState for available users.
	AvailableState UserState = "available"
	// AdminState for admin users.
	AdminState UserState = "admin"
	// ModeratorState for moderator users.
	ModeratorState UserState = "moderator"
	// AdminOrModeratorState for admin or moderator users.
	AdminOrModeratorState UserState = "adminOrModerator"
	// SilencedState for silenced users.
	SilencedState UserState = "silenced"
	// SuspendedState for suspended users.
	SuspendedState UserState = "suspended"
)

// Service is the base for all the endpoints on this service.
type Service struct {
	Call core.RequestHandlerFunc
}

// NewService creates a new Service instance.
func NewService(requestHandler core.RequestHandlerFunc) *Service {
	return &Service{Call: requestHandler}
}
