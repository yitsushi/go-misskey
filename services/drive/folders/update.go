package folders

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// UpdateRequest is the request object for an Update request.
type UpdateRequest struct {
	FolderID string      `json:"folderId"`
	Name     string      `json:"name"`
	ParentID core.String `json:"parentId"`
}

// Validate the request.
func (r UpdateRequest) Validate() error {
	return nil
}

// Update a folder.
//
// The request uses only the ID, Name and ParentID values.
//
// Warning: both Name and ParentID will be updated, so
// if you want to change only the name, fill in the ParentID
// with its current value, if you leave it as null, it will
// be moved to the root of the drive.
func (s *Service) Update(folder models.Folder) (models.Folder, error) {
	request := UpdateRequest{
		FolderID: folder.ID,
		Name:     folder.Name,
		ParentID: folder.ParentID,
	}

	var response models.Folder
	err := s.Call(
		&core.JSONRequest{Request: &request, Path: "/drive/folders/update"},
		&response,
	)

	return response, err
}
