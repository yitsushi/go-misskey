package files

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// ShowRequest is the request to show a file.
// FileID has higher priority, therefore if it's not empty, the endpoint
// will ignore the URL paramater and tries to find a file with given ID.
// As a side-effect, if both of them are defined, but the FileID is different
// than the URL, but both exists, it will return the File with given ID, as it
// simply ignores the URL parameter in this case.
type ShowRequest struct {
	FileID string `json:"fileId"`
	URL    string `json:"url"`
}

// Validate the request.
func (r ShowRequest) Validate() error {
	return nil
}

// Show gets a folder by its ID.
func (s *Service) Show(request ShowRequest) (models.File, error) {
	var response models.File
	err := s.Call(
		&core.JSONRequest{Request: &request, Path: "/drive/files/show"},
		&response,
	)

	return response, err
}
