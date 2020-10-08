package meta

import "github.com/yitsushi/go-misskey/core"

// AnnouncementsRequest represents an Announcement request.
type AnnouncementsRequest struct {
	*core.BaseRequest
	WithUnreads bool   `json:"withUnreads"`
	SinceID     string `json:"sinceId"`
	UntilID     string `json:"untilId"`
}

// AnnouncementsResponse represents an Announcement response.
type AnnouncementsResponse []Announcement

// Announcement represents one announcement.
type Announcement struct {
	ID        core.String `json:"id"`
	CreatedAt core.String `json:"createdAt"`
	UpdatedAt core.String `json:"updatedAt"`
	Title     core.String `json:"title"`
	Text      core.String `json:"text"`
	ImageURL  core.String `json:"imageUrl"`
}
