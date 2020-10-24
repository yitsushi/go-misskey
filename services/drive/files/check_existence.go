package files

import (
	"github.com/yitsushi/go-misskey/core"
)

// CheckExistenceRequest is the request to check
// if a given file with md5 hash exists or not.
type CheckExistenceRequest struct {
	MD5 string `json:"md5"`
}

// Validate the request.
func (r CheckExistenceRequest) Validate() error {
	return nil
}

// CheckExistence check if a file exists or not with given md5.
// md5 hash of the file, not its name.
func (s *Service) CheckExistence(md5 string) (bool, error) {
	var response bool
	err := s.Call(
		&core.JSONRequest{
			Request: &CheckExistenceRequest{
				MD5: md5,
			},
			Path: "/drive/files/check-existence",
		},
		&response,
	)

	return response, err
}
