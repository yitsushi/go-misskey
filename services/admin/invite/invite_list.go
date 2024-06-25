package invite

import (
	"fmt"

	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

type (
	// StatusType defines invite status type such as "unused".
	StatusType string

	// SortOrder defines sort order type of invites.
	SortOrder string
)

const (
	// StatusTypeUnused represents a invite code status is unused.
	StatusTypeUnused StatusType = "unused"

	// StatusTypeUsed represents a invite code status is used.
	StatusTypeUsed StatusType = "used"

	// StatusTypeExpired represents a invite code status is not used and expired.
	StatusTypeExpired StatusType = "expired"

	// StatusTypeAll represents any invite code statuses.
	StatusTypeAll StatusType = "all"

	// SortOrderCreatedAtAsc represents invite codes order by createdAt ascending order.
	SortOrderCreatedAtAsc SortOrder = "-createdAt"

	// SortOrderCreatedAtDesc represents invite codes order by createdAt descending order.
	SortOrderCreatedAtDesc SortOrder = "+createdAt"

	// SortOrderUsedAtAsc represents invite codes order by usedAt ascending order.
	SortOrderUsedAtAsc SortOrder = "-usedAt"

	// SortOrderUsedAtDesc represents invite codes order by usedAt descending order.
	SortOrderUsedAtDesc SortOrder = "+usedAt"
)

// ListRequest is a request object for list invitation codes.
type ListRequest struct {
	Limit  int        `json:"limit"`
	Offset int        `json:"offset"`
	Type   StatusType `json:"type"`
	Sort   SortOrder  `json:"sort"`
}

// Validate the request parameters.
func (r *ListRequest) Validate() error {
	const (
		maxLimit = 100
		minLimit = 1
	)

	if r.Limit > maxLimit || r.Limit < minLimit {
		return core.RequestValidationError{
			Request: r,
			Message: fmt.Sprintf(core.OutOfRangeError, maxLimit, minLimit),
			Field:   "Limit",
		}
	}

	switch r.Type {
	case StatusTypeUnused,
		StatusTypeUsed,
		StatusTypeExpired,
		StatusTypeAll:
		// validation ok
	default:
		return core.RequestValidationError{
			Request: r,
			Message: fmt.Sprintf("unknown status type: %s", r.Type),
			Field:   "Type",
		}
	}

	switch r.Sort {
	case SortOrderCreatedAtAsc,
		SortOrderCreatedAtDesc,
		SortOrderUsedAtAsc,
		SortOrderUsedAtDesc:
	// validation ok
	default:
		return core.RequestValidationError{
			Request: r,
			Message: fmt.Sprintf("unknown sort order: %s", r.Sort),
			Field:   "Sort",
		}
	}

	return nil
}

// List gets list of invites.
// This can narrow down by invite status type and can sort by createdAt or usedAt time.
// limit should be 1 - 100.
func (s *Service) List(limit, offset int, statusType StatusType, sort SortOrder) ([]*models.Invite, error) {
	var response []*models.Invite
	err := s.Call(
		&core.JSONRequest{
			Request: &ListRequest{
				Limit:  limit,
				Offset: offset,
				Type:   statusType,
				Sort:   sort,
			},
			Path: "/admin/invite/list",
		},
		&response,
	)

	return response, err
}
