package drive

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/services/drive/files"
	"github.com/yitsushi/go-misskey/services/drive/folders"
)

const (
	// defaultListLimit is the default value for the limit lists.
	defaultListLimit = 10
)

// Service is the base for all the endpoints on this service.
type Service struct {
	Call core.RequestHandlerFunc
}

// NewService creates a new Service instance.
func NewService(requestHandler core.RequestHandlerFunc) *Service {
	return &Service{Call: requestHandler}
}

// File contains all endpoints under /drive/files.
func (s *Service) File() *files.Service {
	return files.NewService(s.Call)
}

// Folder contains all endpoints under /drive/folders.
func (s *Service) Folder() *folders.Service {
	return folders.NewService(s.Call)
}
