package antennas

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// UpdateRequest is the request object for an Update request.
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

// Validate the request.
func (r UpdateRequest) Validate() error {
	return nil
}

// Update is the endpoint to update an Antenna.
func (s *Service) Update(request UpdateRequest) (models.Antenna, error) {
	var response models.Antenna
	err := s.Call(
		&core.JSONRequest{Request: &request, Path: "/antennas/update"},
		&response,
	)

	return response, err
}

// UpdateAntenna updates an antenna from struct.
func (s *Service) UpdateAntenna(antenna *models.Antenna) (models.Antenna, error) {
	return s.Update(UpdateRequest{
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
