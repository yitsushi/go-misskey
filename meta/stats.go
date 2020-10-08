package meta

import "github.com/yitsushi/go-misskey/core"

// StatsRequest represents an Stats request.
type StatsRequest struct {
	*core.BaseRequest
}

// StatsResponse represents one announcement.
type StatsResponse struct {
	NotesCount         uint64 `json:"notesCount"`
	OriginalNotesCount uint64 `json:"originalNotesCount"`
	UsersCount         uint64 `json:"usersCount"`
	OriginalUsersCount uint64 `json:"originalUsersCount"`
	Instances          uint64 `json:"instances"`
}
