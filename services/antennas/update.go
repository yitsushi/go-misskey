package antennas

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// UpdateRequest is the request object for a Update request.
type UpdateRequest struct {
	AntennaID       string               `json:"antennaId"`
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

// UpdateOptions contains all values that can be used to update an Antenna.
type UpdateOptions struct {
	AntennaID       string
	Name            string
	Source          models.AntennaSource
	UserListID      core.String
	UserGroupID     core.String
	Keywords        [][]string
	ExcludeKeywords [][]string
	Users           []string
	CaseSensitive   bool
	WithReplies     bool
	WithOnlyFile    bool
	Notify          bool
}

// Update is the endpoint to update an Antenna.
func (s *Service) Update(options *UpdateOptions) (models.Antenna, error) {
	request := &UpdateRequest{
		AntennaID:       options.AntennaID,
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
		&core.BaseRequest{Request: request, Path: "/antennas/update"},
		&response,
	)

	return response, err
}

// UpdateAntenna updates an antenna from struct.
func (s *Service) UpdateAntenna(antenna *models.Antenna) (models.Antenna, error) {
	return s.Update(&UpdateOptions{
		AntennaID:       antenna.ID,
		Name:            antenna.Name,
		Source:          antenna.Source,
		UserListID:      antenna.UserListID,
		UserGroupID:     antenna.UserGroupID,
		Keywords:        antenna.Keywords,
		ExcludeKeywords: antenna.ExcludeKeywords,
		Users:           antenna.Users,
		CaseSensitive:   antenna.CaseSensitive,
		WithReplies:     antenna.WithReplies,
		WithOnlyFile:    antenna.WithOnlyFile,
		Notify:          antenna.Notify,
	})
}
