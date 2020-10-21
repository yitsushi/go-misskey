package drive

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// FoldersRequest gets a list of folders available in drive.
type FoldersRequest struct {
	Limit    uint        `json:"limit"`
	SinceID  string      `json:"sinceId"`
	UntilID  string      `json:"untilId"`
	FolderID core.String `json:"folderId"`
}

// FoldersOptions holds all values that can be passed as a parameter for FolderRequest.
type FoldersOptions struct {
	Limit    uint
	SinceID  string
	UntilID  string
	FolderID core.String
}

// Folders lists all folders in drive.
func (s *Service) Folders(options *FoldersOptions) ([]models.Folder, error) {
	request := &FoldersRequest{
		Limit:    options.Limit,
		SinceID:  options.SinceID,
		UntilID:  options.UntilID,
		FolderID: options.FolderID,
	}

	if request.Limit < 1 {
		request.Limit = DefaultListLimit
	}

	var response []models.Folder
	err := s.Call(
		&core.BaseRequest{Request: request, Path: "/drive/folders"},
		&response,
	)

	return response, err
}
