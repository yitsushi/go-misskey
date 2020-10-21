package models

import (
	"time"

	"github.com/yitsushi/go-misskey/core"
)

// Folder is the representation of a Folder entity.
//
// Actually it seems, even if the API documentation says
// it returns with the number of sub-folders and files,
// it doesn't return with them. I'll leave them there,
// as they may appear in the response later, but keep
// in mind, they will be always 0.
type Folder struct {
	ID           string      `json:"id"`
	ParentID     core.String `json:"parentId"`
	CreatedAt    time.Time   `json:"createdAt"`
	Name         string      `json:"name"`
	FoldersCount uint64      `json:"foldersCount"`
	FilesCount   uint64      `json:"filesCount"`
	Parent       *Folder     `json:"parent"`
}
