package antennas

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
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

// Validate the request.
func (r NotesRequest) Validate() error {
	if r.Limit < 1 || r.Limit > 100 {
		return core.RequestValidationError{
			Request: r,
			Message: core.NewRangeError(1, 100),
			Field:   "Limit",
		}
	}

	return nil
}

// Notes is the endpoint to get Notes for an Antenna.
func (s *Service) Notes(request NotesRequest) ([]models.Note, error) {
	var response []models.Note
	err := s.Call(
		&core.JSONRequest{Request: &request, Path: "/antennas/notes"},
		&response,
	)

	return response, err
}
