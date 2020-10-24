package antennas

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// CreateRequest represents an antennas/create request.
type CreateRequest struct {
	Name            string               `json:"name"`
	Source          models.AntennaSource `json:"src"`
	UserListID      core.String          `json:"userListId"`
	UserGroupID     core.String          `json:"userGroupId"`
	Keywords        [][]string           `json:"keywords"`
	ExcludeKeywords [][]string           `json:"excludeKeywords"`
	Users           []string             `json:"users"`
	CaseSensitive   bool                 `json:"caseSensitive"`
	WithReplies     bool                 `json:"withReplies"`
	WithOnlyFile    bool                 `json:"withFile"`
	Notify          bool                 `json:"notify"`
}

// Validate the request.
func (r CreateRequest) Validate() error {
	if r.Name == "" {
		return core.RequestValidationError{
			Request: r,
			Message: core.UndefinedRequiredField,
			Field:   "Name",
		}
	}

	if r.Source == "" {
		return core.RequestValidationError{
			Request: r,
			Message: core.UndefinedRequiredField,
			Field:   "Source",
		}
	}

	if r.Keywords == nil {
		return core.RequestValidationError{
			Request: r,
			Message: core.UndefinedRequiredField,
			Field:   "Keywords",
		}
	}

	if r.ExcludeKeywords == nil {
		return core.RequestValidationError{
			Request: r,
			Message: core.UndefinedRequiredField,
			Field:   "ExcludeKeywords",
		}
	}

	if r.Users == nil {
		return core.RequestValidationError{
			Request: r,
			Message: core.UndefinedRequiredField,
			Field:   "Users",
		}
	}

	return nil
}

// Create antenna endpoint.
func (s *Service) Create(request CreateRequest) (models.Antenna, error) {
	var response models.Antenna
	err := s.Call(
		&core.JSONRequest{Request: &request, Path: "/antennas/create"},
		&response,
	)

	return response, err
}
