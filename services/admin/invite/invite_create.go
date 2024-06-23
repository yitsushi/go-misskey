package invite

import (
	"fmt"
	"time"

	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// CreateRequest is a request object for create a new invitation code.
// This can create multiple invitations at once.
type CreateRequest struct {
	Count     int    `json:"count"`
	ExpiresAt string `json:"expiresAt,omitempty"`
}

// Validate the request parameters.
func (r *CreateRequest) Validate() error {
	const (
		maxCount = 100
		minCount = 1
	)

	if r.Count > maxCount || r.Count < minCount {
		return core.RequestValidationError{
			Request: r,
			Message: fmt.Sprintf(core.OutOfRangeError, maxCount, minCount),
			Field:   "Count",
		}
	}

	if r.ExpiresAt != "" {
		if _, err := time.Parse(time.RFC3339, r.ExpiresAt); err != nil {
			return core.RequestValidationError{
				Request: r,
				Message: fmt.Sprintf("Invalid time format: %s", r.ExpiresAt),
				Field:   "ExpiresAt",
			}
		}
	}

	return nil
}

// Create creates a given number of invites with expired time and returns created invites when request succeeded.
// count should be 1 - 100.
func (s *Service) Create(count int, expiresAt time.Time) ([]*models.Invite, error) {
	var expiresAtStr string
	if !expiresAt.IsZero() {
		expiresAtStr = expiresAt.Format(time.RFC3339)
	}

	var response []*models.Invite
	err := s.Call(
		&core.JSONRequest{
			Request: &CreateRequest{
				Count:     count,
				ExpiresAt: expiresAtStr,
			},
			Path: "/admin/invite/create",
		},
		&response,
	)

	return response, err
}
