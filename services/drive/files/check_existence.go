package files

import (
	"github.com/yitsushi/go-misskey/core"
)

// CheckExistenceRequest is the request to check
// if a given file with md5 hash exists or not.
type CheckExistenceRequest struct {
	MD5 string `json:"md5"`
}

// CheckExistence check if a file exists or not with given md5.
// md5 hash of the file, not its name.
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
