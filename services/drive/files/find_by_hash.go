package files

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// FindByHashRequest is the request to find file(s) with md5.
type FindByHashRequest struct {
	MD5 string `json:"md5"`
}

// FindByHash gets file(s) by their md5 hash.
// If there is no file with given md5 hash, it returns with
// an empty list without error.
func (s *Service) FindByHash(md5 string) ([]models.File, error) {
	request := &FindByHashRequest{
		MD5: md5,
	}

	var response []models.File
	err := s.Call(
		&core.BaseRequest{Request: request, Path: "/drive/files/find-by-hash"},
		&response,
	)

	return response, err
}
