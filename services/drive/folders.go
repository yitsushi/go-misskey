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

// Validate the request.
func (r FoldersRequest) Validate() error {
	if r.Limit < 1 || r.Limit > 100 {
		return core.RequestValidationError{
			Request: r,
			Message: core.NewRangeError(1, 100),
			Field:   "Limit",
		}
	}

	return nil
}

// Folders lists all folders in drive.
func (s *Service) Folders(request FoldersRequest) ([]models.Folder, error) {
	var response []models.Folder
	err := s.Call(
		&core.JSONRequest{Request: &request, Path: "/drive/folders"},
		&response,
	)

	return response, err
}
