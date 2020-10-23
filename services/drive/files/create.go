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

// CreateOptions has all the values you can play with.
type CreateOptions struct {
	FolderID    string
	Name        string
	IsSensitive bool
	Force       bool
	Content     []byte
}

// Create a file.
func (s *Service) Create(options *CreateOptions) (models.File, error) {
	request := CreateRequest{
		FolderID:    options.FolderID,
		Name:        options.Name,
		IsSensitive: options.IsSensitive,
		Force:       options.Force,
		Content:     options.Content,
	}

	var response models.File
	err := s.Call(
		&core.MultipartRequest{Request: request, Path: "/drive/files/create"},
		&response,
	)

	return response, err
}
