package instance

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// ServerInfoRequest represents a Clear request.
type GetTableStatsRequest struct{}

// Validate the request.
func (r GetTableStatsRequest) Validate() error {
	return nil
}

// GetTableStats gets the table statistics.
func (s *Service) GetTableStats() (models.TableStats, error) {
	var response models.TableStats
	err := s.Call(
		&core.JSONRequest{Request: &GetTableStatsRequest{}, Path: "/admin/get-table-stats"},
		&response,
	)

	return response, err
}
