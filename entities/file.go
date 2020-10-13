package entities

import (
	"time"

	"github.com/yitsushi/go-misskey/core"
)

// File is a file.
type File struct {
	ID           string         `json:"id"`
	CreatedAt    time.Time      `json:"createdAt"`
	Name         core.String    `json:"name"`
	Type         string         `json:"type"`
	Md5          string         `json:"md5"`
	Size         uint64         `json:"size"`
	IsSensitive  bool           `json:"isSensitive"`
	Blurhash     string         `json:"blurhash"`
	Properties   FileProperties `json:"properties"`
	URL          string         `json:"url"`
	ThumbnailURL string         `json:"thumbnailUrl"`
	FolderID     core.String    `json:"folderId"`
	Folder       interface{}    `json:"folder"`
	User         interface{}    `json:"user"`
}

// FileProperties holds the properties of the file like width and height.
// I don't know if there are other too, so I'll go that way now.
type FileProperties struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}
