package antennas

import "github.com/yitsushi/go-misskey/core"

type AntennaSource string

const (
	HomeSrc  AntennaSource = "home"
	AllSrc   AntennaSource = "all"
	UsersSrc AntennaSource = "users"
	ListSrc  AntennaSource = "list"
	GroupSrc AntennaSource = "group"
)

// CreateRequest represents an antennas/create request.
type CreateRequest struct {
	*core.BaseRequest
	Name            string        `json:"name"`
	Source          AntennaSource `json:"src"`
	UserListID      core.String   `json:"userListId"`
	UserGroupId     core.String   `json:"userGroupId"`
	Keywords        [][]string    `json:"keywords"`
	ExcludeKeywords [][]string    `json:"excludeKeywords"`
	Users           []string      `json:"users"`
	CaseSensitive   bool          `json:"caseSensitive"`
	WithReplies     bool          `json:"withReplies"`
	WithOnlyFile    bool          `json:"withFile"`
	Notify          bool          `json:"notify"`
}

// CreateResponse represents the response on an antennas/create request.
type CreateResponse struct {
	*Antenna
}

// CreateOptions contains all values that can be used to create an Antenna.
type CreateOptions struct {
	Name        string
	Source      AntennaSource
	UserListID  core.String
	UserGroupId core.String
	// Keywords is a simple array of strings,
	// but in the background it's an array of arrays.
	// The outer array has an OR condition,
	// the inner one has AND condition.
	// For simplicity it's just array of strings where
	// items has OR condition, for AND condition just put
	// all items in one space separated string.
	Keywords []string
	// ExcludeKeywords is a simple array of strings,
	// but in the background it's an array of arrays.
	// The outer array has an OR condition,
	// the inner one has AND condition.
	// For simplicity it's just array of strings where
	// items has OR condition, for AND condition just put
	// all items in one space separated string.
	ExcludeKeywords []string
	Users           []string
	CaseSensitive   bool
	WithReplies     bool
	WithOnlyFile    bool
	Notify          bool
}

func (s *Service) Create(options *CreateOptions) (CreateResponse, error) {
	request := &CreateRequest{
		Name:            options.Name,
		Source:          options.Source,
		UserListID:      options.UserListID,
		UserGroupId:     options.UserGroupId,
		Keywords:        [][]string{options.Keywords, options.Keywords},
		ExcludeKeywords: [][]string{options.ExcludeKeywords},
		Users:           options.Users,
		CaseSensitive:   options.CaseSensitive,
		WithReplies:     options.WithReplies,
		WithOnlyFile:    options.WithOnlyFile,
		Notify:          options.Notify,
	}

	var response CreateResponse
	err := s.Call(
		&core.BaseRequest{Request: request, Path: "/antennas/create"},
		&response,
	)

	return response, err
}
