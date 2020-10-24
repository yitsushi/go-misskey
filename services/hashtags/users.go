package hashtags

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// UserState of user.
type UserState string

// UserOrigin of a user.
type UserOrigin string

const (
	// StateAll for all users.
	StateAll UserState = "all"
	// StateAlive for only alive users.
	StateAlive UserState = "alive"

	// OriginCombined for local and remote users.
	OriginCombined UserOrigin = "combined"
	// OriginLocal for only localusers.
	OriginLocal UserOrigin = "local"
	// OriginRemote for only remote users.
	OriginRemote UserOrigin = "remote"

	// DefaultState is the default state.
	DefaultState UserState = StateAll
	// DefaultOrigin is the default origin.
	DefaultOrigin UserOrigin = OriginLocal
)

// UsersRequest represents an Users request.
type UsersRequest struct {
	Tag    string     `json:"tag"`
	Limit  uint       `json:"limit"`
	Sort   string     `json:"sort"`
	State  UserState  `json:"state"`
	Origin UserOrigin `json:"origin"`
}

// UsersOptions are all the options you can play with.
type UsersOptions struct {
	Tag    string
	Limit  uint
	Sort   string
	State  UserState
	Origin UserOrigin
}

// Validate options.
func (options *UsersOptions) Validate() error {
	missingError := core.MissingOptionsError{
		Endpoint:      "Hashtags/Users",
		Struct:        "UsersOptions",
		MissingFields: []string{},
	}

	if options.Sort == "" {
		missingError.MissingFields = append(missingError.MissingFields, "Sort")
	}

	if options.Tag == "" {
		missingError.MissingFields = append(missingError.MissingFields, "Tag")
	}

	if len(missingError.MissingFields) == 0 {
		return nil
	}

	return missingError
}

// Users endpoint.
func (s *Service) Users(options *UsersOptions) ([]models.User, error) {
	var response []models.User

	if options == nil {
		return response, core.MissingOptionsError{
			Endpoint: "Hashtags/Users",
			Struct:   "UsersOptions",
		}
	}

	err := options.Validate()
	if err != nil {
		return response, err
	}

	if options.State == "" {
		options.State = DefaultState
	}

	if options.Origin == "" {
		options.Origin = DefaultOrigin
	}

	request := UsersRequest{
		Limit:  options.Limit,
		Sort:   options.Sort,
		Tag:    options.Tag,
		State:  options.State,
		Origin: options.Origin,
	}

	err = s.Call(
		&core.BaseRequest{Request: &request, Path: "/hashtags/users"},
		&response,
	)

	return response, err
}
