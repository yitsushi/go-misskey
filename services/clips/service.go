package clips

import "github.com/yitsushi/go-misskey/core"

const (
	// maximumNameLength is the maximum length of the name of the clip.
	maximumNameLength = 100
)

// Service is the base for all the endpoints on this service.
type Service struct {
	Call core.RequestHandlerFunc
}

// NewService creates a new Service instance.
func NewService(requestHandler core.RequestHandlerFunc) *Service {
	return &Service{Call: requestHandler}
}
