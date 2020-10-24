package hashtags

import (
	"fmt"

	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// SortFlag is used for sorting on the List endpoint.
// Here there are two functions on it, Ascending and Descending,
// and they specify how the tags should be sorted.
//
// In the background, the endpoint expect a string
// with a + or a - sign.
type SortFlag string

const (
	// SortMentionedUsers sorts hashtags by user mentions.
	SortMentionedUsers SortFlag = "mentionedUsers"
	// SortMentionedLocalUsers sorts hashtags by local only user mentions.
	SortMentionedLocalUsers SortFlag = "mentionedLocalUsers"
	// SortMentionedRemoteUsers sorts hashtags by remote only user mentions.
	SortMentionedRemoteUsers SortFlag = "mentionedRemoteUsers"
	// SortAttachedUsers sorts hashtags by user attachment.
	SortAttachedUsers SortFlag = "attachedUsers"
	// SortAttachedLocalUsers sorts hashtags by local only user attachment.
	SortAttachedLocalUsers SortFlag = "attachedLocalUsers"
	// SortAttachedRemoteUsers sorts hashtags by remote only user attachment.
	SortAttachedRemoteUsers SortFlag = "attachedRemoteUsers"
)

// Descending order.
func (s SortFlag) Descending() string {
	return fmt.Sprintf("-%s", s)
}

// Ascending order.
func (s SortFlag) Ascending() string {
	return fmt.Sprintf("+%s", s)
}

// ListRequest represents an List request.
type ListRequest struct {
	Limit                    uint   `json:"limit"`
	AttachedToUserOnly       bool   `json:"attachedToUserOnly"`
	AttachedToLocalUserOnly  bool   `json:"attachedToLocalUserOnly"`
	AttachedToRemoteUserOnly bool   `json:"attachedToRemoteUserOnly"`
	Sort                     string `json:"sort"`
}

// ListOptions are all the options you can play with.
type ListOptions struct {
	Limit                    uint
	AttachedToUserOnly       bool
	AttachedToLocalUserOnly  bool
	AttachedToRemoteUserOnly bool
	Sort                     string
}

// List endpoint.
func (s *Service) List(options *ListOptions) ([]models.Hashtag, error) {
	var response []models.Hashtag

	if options == nil {
		return response, core.MissingOptionsError{
			Endpoint: "Hashtags/List",
			Struct:   "ListOptions",
		}
	}

	if options.Sort == "" {
		return response, core.MissingOptionsError{
			Endpoint:      "Hashtags/List",
			Struct:        "ListOptions",
			MissingFields: []string{"Sort"},
		}
	}

	request := ListRequest{
		Limit:                    options.Limit,
		AttachedToUserOnly:       options.AttachedToUserOnly,
		AttachedToLocalUserOnly:  options.AttachedToLocalUserOnly,
		AttachedToRemoteUserOnly: options.AttachedToRemoteUserOnly,
		Sort:                     options.Sort,
	}

	err := s.Call(
		&core.BaseRequest{Request: &request, Path: "/hashtags/list"},
		&response,
	)

	return response, err
}
