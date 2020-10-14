package models

import (
	"time"

	"github.com/yitsushi/go-misskey/core"
)

// AntennaSource is just an "enum" like type alias.
type AntennaSource string

const (
	// HomeSrc is "home" as source for an Antenna.
	HomeSrc AntennaSource = "home"
	// AllSrc is "all" as source for an Antenna.
	AllSrc AntennaSource = "all"
	// UsersSrc is "users" as source for an Antenna.
	UsersSrc AntennaSource = "users"
	// ListSrc is "list" as source for an Antenna.
	ListSrc AntennaSource = "list"
	// GroupSrc is "group" as source for an Antenna.
	GroupSrc AntennaSource = "group"
)

// Antenna is a pure representation of an Antenna.
type Antenna struct {
	ID              string        `json:"id"`
	CreatedAt       time.Time     `json:"createdAt"`
	Name            string        `json:"name"`
	Keywords        [][]string    `json:"keywords"`
	ExcludeKeywords [][]string    `json:"excludeKeywords"`
	Source          AntennaSource `json:"src"`
	UserListID      core.String   `json:"userListId"`
	UserGroupID     core.String   `json:"userGroupId"`
	Users           []string      `json:"users"`
	CaseSensitive   bool          `json:"caseSensitive"`
	Notify          bool          `json:"notify"`
	WithReplies     bool          `json:"withReplies"`
	WithOnlyFile    bool          `json:"withFile"`
	HasUnreadNote   bool          `json:"hasUnreadNote"`
}
