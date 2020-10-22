package folders

import "github.com/yitsushi/go-misskey/core"

// Service is the service for all the endpoints under /drive/folders/*.
type Service struct {
	Call core.RequestHandlerFunc
}

// NewService creates a new Service instance.
func NewService(requestHandler core.RequestHandlerFunc) *Service {
	return &Service{Call: requestHandler}
}
