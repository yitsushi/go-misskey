package users

import (
	"regexp"

	"github.com/yitsushi/go-misskey/core"
)

// AvailableRequest represents an Available request.
type AvailableRequest struct {
	Username string `json:"username"`
}

// AvailableResponse represents an Available response.
type AvailableResponse struct {
	Available bool `json:"available"`
}

// Validate the request.
func (r AvailableRequest) Validate() error {
	// #/properties/username/pattern -> must match pattern "^\w{1,20}$"
	isValid := regexp.MustCompile(`^\w{1,20}$`).MatchString(r.Username)
	if !isValid {
		return core.RequestValidationError{
			Request: r,
			Message: `must match pattern "^\w{1,20}$"`,
			Field:   "#/properties/username/pattern",
		}
	}

	return nil
}

// IsUsernameAvailable is used to determine if a given username is available.
// It does not require authentication.
// It returns an error if the username is invalid.
// The response will be true if the username is available; false if not.
func (s *Service) IsUsernameAvailable(username string) (bool, error) {
	var response AvailableResponse

	request := AvailableRequest{Username: username}

	err := s.Call(
		&core.JSONRequest{Request: &request, Path: "/username/available"},
		&response,
	)
	if err != nil {
		return false, err
	}

	return response.Available, nil
}
