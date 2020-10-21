package models

import (
	"time"

	"github.com/yitsushi/go-misskey/core"
)

// Folder is the representation of a Folder entity.
//
// Actually it seems, even if the API documentation says
// it returns with the number of sub-folders and files,
// it doesn't return always with them. Some of the endpoints return
// with the values like Show, but mostly not.
type Folder struct {
	ID           string      `json:"id"`
	ParentID     core.String `json:"parentId"`
	CreatedAt    time.Time   `json:"createdAt"`
	Name         string      `json:"name"`
	FoldersCount uint64      `json:"foldersCount"`
	FilesCount   uint64      `json:"filesCount"`
	Parent       *Folder     `json:"parent"`
}
