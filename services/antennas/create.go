package antennas

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// CreateRequest represents an antennas/create request.
type CreateRequest struct {
	Name            string               `json:"name"`
	Source          models.AntennaSource `json:"src"`
	UserListID      core.String          `json:"userListId"`
	UserGroupID     core.String          `json:"userGroupId"`
	Keywords        [][]string           `json:"keywords"`
	ExcludeKeywords [][]string           `json:"excludeKeywords"`
	Users           []string             `json:"users"`
	CaseSensitive   bool                 `json:"caseSensitive"`
	WithReplies     bool                 `json:"withReplies"`
	WithOnlyFile    bool                 `json:"withFile"`
	Notify          bool                 `json:"notify"`
}

// CreateOptions contains all values that can be used to create an Antenna.
type CreateOptions struct {
	Name        string
	Source      models.AntennaSource
	UserListID  core.String
	UserGroupID core.String
	// The outer array has an OR condition,
	// the inner one has AND condition.
	Keywords [][]string
	// The outer array has an OR condition,
	// the inner one has AND condition.
	ExcludeKeywords [][]string
	Users           []string
	CaseSensitive   bool
	WithReplies     bool
	WithOnlyFile    bool
	Notify          bool
}

// Create antenna endpoint.
func (s *Service) Create(options *CreateOptions) (models.Antenna, error) {
	request := &CreateRequest{
		Name:            options.Name,
		Source:          options.Source,
		UserListID:      options.UserListID,
		UserGroupID:     options.UserGroupID,
		Keywords:        options.Keywords,
		ExcludeKeywords: options.ExcludeKeywords,
		Users:           options.Users,
		CaseSensitive:   options.CaseSensitive,
		WithReplies:     options.WithReplies,
		WithOnlyFile:    options.WithOnlyFile,
		Notify:          options.Notify,
	}

	var response models.Antenna
	err := s.Call(
		&core.BaseRequest{Request: request, Path: "/antennas/create"},
		&response,
	)

	return response, err
}
