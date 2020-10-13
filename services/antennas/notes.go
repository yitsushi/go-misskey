package antennas

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/entities"
)

const (
	// NoteListDefaultLimit is the default value for notes list.
	NoteListDefaultLimit = 10
)

// NotesRequest represents a request to fetch notes for a given Antenna.
type NotesRequest struct {
	AntennaID string `json:"antennaId"`
	Limit     uint64 `json:"limit"`
	SinceID   string `json:"sinceId"`
	UntilID   string `json:"untilId"`
}

// NotesResponse represents the response for note list.
type NotesResponse []entities.Note

// NotesOptions are all the options available for a Notes request.
type NotesOptions struct {
	AntennaID string
	Limit     uint64
	SinceID   string
	UntilID   string
}

// Notes is the endpoint to get Notes for an Antenna.
func (s *Service) Notes(options *NotesOptions) (NotesResponse, error) {
	request := &NotesRequest{
		AntennaID: options.AntennaID,
		Limit:     options.Limit,
		SinceID:   options.SinceID,
		UntilID:   options.UntilID,
	}

	if request.Limit < 1 || options.Limit > 100 {
		request.Limit = NoteListDefaultLimit
	}

	var response NotesResponse
	err := s.Call(
		&core.BaseRequest{Request: request, Path: "/antennas/notes"},
		&response,
	)

	return response, err
}
