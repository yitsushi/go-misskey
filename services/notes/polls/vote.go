package polls

import (
	"github.com/yitsushi/go-misskey/core"
)

// VoteRequest represents an Vote request.
type VoteRequest struct {
	NoteID string `json:"noteId"`
	Choice int    `json:"choice"`
}

// Validate the request.
func (r VoteRequest) Validate() error {
	return nil
}

// Vote endpoint.
func (s *Service) Vote(noteID string, choice int) error {
	return s.Call(
		&core.JSONRequest{
			Request: &VoteRequest{
				NoteID: noteID,
				Choice: choice,
			},
			Path: "/notes/polls/vote",
		},
		&core.EmptyResponse{},
	)
}
