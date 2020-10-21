package files

import (
	"github.com/yitsushi/go-misskey/core"
)

// CheckExistenceRequest list all notes where a given file has reference.
type CheckExistenceRequest struct {
	MD5 string `json:"md5"`
}

// CheckExistence gets drive information.
func (s *Service) CheckExistence(md5 string) (bool, error) {
	request := &CheckExistenceRequest{
		MD5: md5,
	}

	var response bool
	err := s.Call(
		&core.BaseRequest{Request: request, Path: "/drive/files/check-existence"},
		&response,
	)

	return response, err
}
