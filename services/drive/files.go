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

// FilesRequestWithoutType is the same as FilesRequest, but without the Type
// field. The reason is simple:
// This field is marked as optional and nullable on the endpoint, but has no default
// value defined like folderID has, therefore, even if it's nullable, if it's defined
// as null, the endpoint throws back an error, that it can't be null.
// Endpoint reference:
// https://github.com/syuilo/misskey/blob/develop/src/server/api/endpoints/drive/files.ts#L39
type FilesRequestWithoutType struct {
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

// Information gets drive information.
func (s *Service) Files(options *FilesOptions) ([]models.File, error) {
	request := &FilesRequest{
		Limit:    options.Limit,
		SinceID:  options.SinceID,
		UntilID:  options.UntilID,
		FolderID: options.FolderID,
		Type:     options.Type,
	}

	var response []models.File
	err := s.Call(
		&core.BaseRequest{Request: request, Path: "/drive/files"},
		&response,
	)

	return response, err
}
