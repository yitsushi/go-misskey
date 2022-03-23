package app

import "github.com/yitsushi/go-misskey/core"

const (
	// MaximumNameLength is the maximum length of the name of the clip.
	MaximumNameLength = 100
)

// Service is the base for all the endpoints on this service.
type Service struct {
	Call core.RequestHandlerFunc
}

// NewService creates a new Service instance.
func NewService(requestHandler core.RequestHandlerFunc) *Service {
	return &Service{Call: requestHandler}
}
