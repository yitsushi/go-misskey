package drive

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// StreamRequest gets a list of files available in drive as a flat list,
// all files are included and no filter on folder.
type StreamRequest struct {
	Limit   uint   `json:"limit"`
	SinceID string `json:"sinceId"`
	UntilID string `json:"untilId"`
	Type    string `json:"type"`
}

// Validate the request.
func (r StreamRequest) Validate() error {
	if r.Limit < 1 || r.Limit > 100 {
		return core.RequestValidationError{
			Request: r,
			Message: core.NewRangeError(1, 100),
			Field:   "Limit",
		}
	}

	return nil
}

// Stream lists all folders in drive.
func (s *Service) Stream(request StreamRequest) ([]models.File, error) {
	var response []models.File
	err := s.Call(
		&core.JSONRequest{Request: &request, Path: "/drive/stream"},
		&response,
	)

	return response, err
}
