package notes

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// CreateRequest represents an Create request.
type CreateRequest struct {
	Visibility        models.Visibility `json:"visibility"`
	VisibleUserIDs    []string          `json:"visibleUserIds,omitempty"`
	Text              core.String       `json:"text,omitempty"`
	CW                core.String       `json:"cw,omitempty"`
	ViaMobile         bool              `json:"viaMobile"`
	LocalOnly         bool              `json:"localOnly"`
	NoExtractMentions bool              `json:"noExtractMentions"`
	NoExtractHashtags bool              `json:"noExtractHashtags"`
	NoExtractEmojis   bool              `json:"noExtractEmojis"`
	FileIDs           []string          `json:"fileIds,omitempty"`
	ReplyID           core.String       `json:"replyId,omitempty"`
	RenoteID          core.String       `json:"renoteId,omitempty"`
	ChannelID         core.String       `json:"channelId,omitempty"`
	Poll              *Poll             `json:"poll,omitempty"`
}

// Poll is a poll in a CreateRequest.
type Poll struct {
	Choices      []string `json:"choices,omitempty"`
	Multiple     bool     `json:"multiple,omitempty"`
	ExpiresAt    uint64   `json:"expiresAt,omitempty"`
	ExpiredAfter uint64   `json:"expiredAfter,omitempty"`
}

// CreateResponse is the response for a Create request.
type CreateResponse struct {
	CreatedNote models.Note `json:"createdNote"`
}

const (
	// MinimumNumberOfChoices is the minimum number
	// of choices in a poll.
	MinimumNumberOfChoices = 2
)

// Validate the request.
func (r CreateRequest) Validate() error {
	if r.Poll != nil {
		if len(r.Poll.Choices) < MinimumNumberOfChoices {
			return core.RequestValidationError{
				Request: r,
				Message: core.UndefinedRequiredField,
				Field:   "Poll.Choices",
			}
		}
	}

	if r.Text != nil && core.StringValue(r.Text) == "" {
		return core.RequestValidationError{
			Request: r,
			Message: core.UndefinedRequiredField,
			Field:   "Text",
		}
	}

	if r.Text == nil && len(r.FileIDs) == 0 && r.Poll == nil {
		return core.RequestValidationError{
			Request: r,
			Message: core.UndefinedRequiredField,
			Field:   "Text || FileIDs || Poll",
		}
	}

	return nil
}

// Create endpoint.
func (s *Service) Create(request CreateRequest) (CreateResponse, error) {
	var response CreateResponse

	err := s.Call(
		&core.JSONRequest{Request: &request, Path: "/notes/create"},
		&response,
	)

	return response, err
}
