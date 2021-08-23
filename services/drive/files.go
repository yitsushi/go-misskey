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

// Validate the request.
func (r FilesRequest) Validate() error {
	if r.Limit < 1 || r.Limit > maxLimit {
		return core.RequestValidationError{
			Request: r,
			Message: core.NewRangeError(1, maxLimit),
			Field:   "Limit",
		}
	}

	return nil
}

// Files lists all files in drive.
func (s *Service) Files(request FilesRequest) ([]models.File, error) {
	var response []models.File
	err := s.Call(
		&core.JSONRequest{Request: &request, Path: "/drive/files"},
		&response,
	)

	return response, err
}
