package models

import (
	"time"

	"github.com/yitsushi/go-misskey/core"
)

// Folder is the representation of a Folder entity.
type Folder struct {
	ID           string      `json:"id"`
	ParentID     core.String `json:"parentId"`
	CreatedAt    time.Time   `json:"createdAt"`
	Name         string      `json:"name"`
	FoldersCount uint64      `json:"foldersCount"`
	FilessCount  uint64      `json:"filesCount"`
	Parent       *Folder     `json:"parent"`
}
