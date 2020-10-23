package folders

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// CreateRequest is the request to create a folder.
type CreateRequest struct {
	Name     string      `json:"name"`
	ParentID core.String `json:"parentId"`
}

// CreateOptions are the possible parameters for a Create request.
type CreateOptions struct {
	Name     string
	ParentID core.String
}

// Create a folder.
func (s *Service) Create(options *CreateOptions) (models.Folder, error) {
	request := &CreateRequest{
		Name:     options.Name,
		ParentID: options.ParentID,
	}

	var response models.Folder
	err := s.Call(
		&core.BaseRequest{Request: request, Path: "/drive/folders/create"},
		&response,
	)

	return response, err
}
