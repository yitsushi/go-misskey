package entities

import (
	"time"

	"github.com/yitsushi/go-misskey/core"
)

// Antenna is a pure representation of an Antenna.
type Antenna struct {
	ID              string      `json:"id"`
	CreatedAt       time.Time   `json:"createdAt"`
	Name            string      `json:"name"`
	Keywords        [][]string  `json:"keywords"`
	ExcludeKeywords [][]string  `json:"excludeKeywords"`
	Src             string      `json:"src"`
	UserListID      core.String `json:"userListId"`
	UserGroupID     core.String `json:"userGroupId"`
	Users           []string    `json:"users"`
	CaseSensitive   bool        `json:"caseSensitive"`
	Notify          bool        `json:"notify"`
	WithReplies     bool        `json:"withReplies"`
	WithOnlyFile    bool        `json:"withFile"`
	HasUnreadNote   bool        `json:"hasUnreadNote"`
}
