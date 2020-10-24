package files

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// FindByHashRequest is the request to find file(s) with md5.
type FindByHashRequest struct {
	MD5 string `json:"md5"`
}

// Validate the request.
func (r FindByHashRequest) Validate() error {
	return nil
}

// FindByHash gets file(s) by their md5 hash.
// If there is no file with given md5 hash, it returns with
// an empty list without error.
func (s *Service) FindByHash(md5 string) ([]models.File, error) {
	var response []models.File
	err := s.Call(
		&core.JSONRequest{
			Request: &FindByHashRequest{
				MD5: md5,
			},
			Path: "/drive/files/find-by-hash",
		},
		&response,
	)

	return response, err
}
