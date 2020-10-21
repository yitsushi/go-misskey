package files

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// UpdateRequest is the request object for an Update request.
type UpdateRequest struct {
	FileID      string      `json:"fileId"`
	Name        core.String `json:"name"`
	FolderID    core.String `json:"folderId"`
	IsSensitive bool        `json:"isSensitive"`
}

// Update a file.
//
// The request uses only the ID, Name and FolderID values.
//
// Warning: Name, FolderID and the isSensitive flag (NSFW) will
// be updated, so if you want to change only the name,
// fill in the FolderID and IsSensitive value with their
// current value, if you leave it as null or false,
// it will be moved to the root of the drive and marked as SFW.
func (s *Service) Update(file models.File) (models.File, error) {
	request := &UpdateRequest{
		FileID:      file.ID,
		Name:        file.Name,
		FolderID:    file.FolderID,
		IsSensitive: file.IsSensitive,
	}

	var response models.File
	err := s.Call(
		&core.BaseRequest{Request: request, Path: "/drive/files/update"},
		&response,
	)

	return response, err
}
