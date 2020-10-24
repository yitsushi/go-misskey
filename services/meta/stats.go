package meta

import (
	"github.com/yitsushi/go-misskey/core"
)

// StatsRequest represents an Stats request.
type StatsRequest struct {
}

// Validate the request.
func (r StatsRequest) Validate() error {
	return nil
}

// StatsResponse represents the response to a stats request.
type StatsResponse struct {
	NotesCount         uint64        `json:"notesCount"`
	OriginalNotesCount uint64        `json:"originalNotesCount"`
	UsersCount         uint64        `json:"usersCount"`
	OriginalUsersCount uint64        `json:"originalUsersCount"`
	Instances          uint64        `json:"instances"`
	DriveUsageLocal    core.DataSize `json:"driveUsageLocal"`
	DriveUsageRemote   core.DataSize `json:"driveUsageRemote"`
}

// Stats endpoint.
func (s *Service) Stats() (StatsResponse, error) {
	var response StatsResponse

	err := s.Call(
		&core.JSONRequest{Request: &StatsRequest{}, Path: "/stats"},
		&response,
	)

	return response, err
}
