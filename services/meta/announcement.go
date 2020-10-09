package meta

import (
	"time"

	"github.com/yitsushi/go-misskey/core"
)

// Announcement represents one announcement.
type Announcement struct {
	ID        core.String `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Title     core.String `json:"title"`
	Text      core.String `json:"text"`
	ImageURL  core.String `json:"imageUrl"`
}
