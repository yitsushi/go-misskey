package drive

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// FilesRequest gets a list of files available in drive.
type FilesRequest struct {
	Limit    uint        `json:"limit"`
	SinceID  string      `json:"sinceId"`
	UntilID  string      `json:"untilId"`
	FolderID core.String `json:"folderId"`
	Type     core.String `json:"type"`
}

// FilesOptions holds all values that can be passed as a parameter for FileRequest.
type FilesOptions struct {
	Limit    uint
	SinceID  string
	UntilID  string
	FolderID core.String
	Type     core.String
}

// Files lists all files in drive.
func (s *Service) Files(options *FilesOptions) ([]models.File, error) {
	request := &FilesRequest{
		Limit:    options.Limit,
		SinceID:  options.SinceID,
		UntilID:  options.UntilID,
		FolderID: options.FolderID,
		Type:     options.Type,
	}

	if request.Limit < 1 {
		request.Limit = DefaultListLimit
	}

	var response []models.File
	err := s.Call(
		&core.BaseRequest{Request: request, Path: "/drive/files"},
		&response,
	)

	return response, err
}
