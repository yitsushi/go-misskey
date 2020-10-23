package folders

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// FindRequest is the request to find folder(s) by their and parent folder.
type FindRequest struct {
	Name     string      `json:"name"`
	ParentID core.String `json:"parentId"`
}

// FindOptions are the possible parameters for a Find request.
type FindOptions struct {
	Name     string
	ParentID core.String
}

// Find gets folder(s) by their name and parent folder.
func (s *Service) Find(options *FindOptions) ([]models.Folder, error) {
	request := &FindRequest{
		Name:     options.Name,
		ParentID: options.ParentID,
	}

	var response []models.Folder
	err := s.Call(
		&core.BaseRequest{Request: request, Path: "/drive/folders/find"},
		&response,
	)

	return response, err
}
