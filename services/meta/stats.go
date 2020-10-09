package meta

import (
	"github.com/yitsushi/go-misskey/core"
)

// StatsRequest represents an Stats request.
type StatsRequest struct {
	*core.BaseRequest
}

// StatsResponse represents the response to a stats request.
type StatsResponse struct {
	NotesCount         uint64 `json:"notesCount"`
	OriginalNotesCount uint64 `json:"originalNotesCount"`
	UsersCount         uint64 `json:"usersCount"`
	OriginalUsersCount uint64 `json:"originalUsersCount"`
	Instances          uint64 `json:"instances"`
}

func (s *Service) Stats() (StatsResponse, error) {
	var response StatsResponse

	err := s.Call(
		&core.BaseRequest{Request: &StatsRequest{}, Path: "/stats"},
		&response,
	)

	return response, err
}
