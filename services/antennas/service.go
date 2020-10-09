package antennas

import "github.com/yitsushi/go-misskey/core"

// Service is the base for all the endpoints on this service.
type Service struct {
	Call core.RequestHandlerFunc
}

var service *Service

// NewService creates a new Service or use an existing one (singleton).
func NewService(requestHandler core.RequestHandlerFunc) *Service {
	if service == nil {
		service = &Service{Call: requestHandler}
	}
	return service
}

