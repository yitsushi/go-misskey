package files

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// CreateRequest represents a request to create a file.
type CreateRequest struct {
	FolderID    string `multipart:"folderId,type=field"`
	Name        string `multipart:"name,type=field"`
	IsSensitive bool   `multipart:"isSensitive,type=field"`
	Force       bool   `multipart:"force,type=field"`
	Content     []byte `multipart:"ref=name,type=file"`
}

// Validate the request.
func (r CreateRequest) Validate() error {
	return nil
}

// Create a file.
func (s *Service) Create(request CreateRequest) (models.File, error) {
	var response models.File
	err := s.Call(
		&core.MultipartRequest{Request: request, Path: "/drive/files/create"},
		&response,
	)

	return response, err
}
