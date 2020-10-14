package models

import "github.com/yitsushi/go-misskey/core"

// Emoji is a simple reprosentation of an Emoji.
type Emoji struct {
	ID       core.String   `json:"id"`
	Aliases  []core.String `json:"aliases"`
	Name     core.String   `json:"name"`
	Category core.String   `json:"category"`
	Host     core.String   `json:"host"`
	URL      core.String   `json:"url"`
}
