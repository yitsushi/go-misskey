package models

import (
	"time"

	"github.com/yitsushi/go-misskey/core"
)

// Note is a note.
type Note struct {
	ID           string            `json:"id"`
	CreatedAt    time.Time         `json:"createdAt"`
	UserID       string            `json:"userId"`
	User         User              `json:"user"`
	Text         string            `json:"text"`
	CW           core.String       `json:"cw"`
	Visibility   string            `json:"visibility"`
	Mentions     []string          `json:"mentions"`
	RenoteCount  uint64            `json:"renoteCount"`
	RepliesCount uint64            `json:"repliesCount"`
	Reactions    map[string]uint64 `json:"reactions"`
	Tags         []string          `json:"tags"`
	Emojis       []Emoji           `json:"emojis"`
	FileIds      []string          `json:"fileIds"`
	Files        []File            `json:"files"`
	Reply        *Note             `json:"reply"`
	ReplyID      core.String       `json:"replyId"`
	RenoteID     core.String       `json:"renoteId"`
	URI          string            `json:"uri"`
	URL          string            `json:"url"`
	Instance     *Instance         `json:"instance"`
}
