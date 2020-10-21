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

// StreamOptions holds all values that can be passed as a parameter for StreamRequest.
type StreamOptions struct {
	Limit   uint
	SinceID string
	UntilID string
	Type    string
}

// Stream lists all folders in drive.
func (s *Service) Stream(options *StreamOptions) ([]models.File, error) {
	request := &StreamRequest{
		Limit:   options.Limit,
		SinceID: options.SinceID,
		UntilID: options.UntilID,
		Type:    options.Type,
	}

	if request.Limit < 1 {
		request.Limit = DefaultListLimit
	}

	var response []models.File
	err := s.Call(
		&core.BaseRequest{Request: request, Path: "/drive/stream"},
		&response,
	)

	return response, err
}
