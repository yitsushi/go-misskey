package models

import (
	"time"

	"github.com/yitsushi/go-misskey/core"
)

// File is a file.
// Folder and User is defined as an interface{} because
// I have no idea how it looks like.
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
	Comment      string         `json:"comment"`
	ThumbnailURL string         `json:"thumbnailUrl"`
	FolderID     core.String    `json:"folderId"`
	Folder       *Folder        `json:"folder"`
	User         *User          `json:"user"`
}

// FileProperties holds the properties of the file like width and height.
// I don't know if there are other too, so I'll go that way now.
type FileProperties struct {
	Width       int    `json:"width"`
	Height      int    `json:"height"`
	Orientation int    `json:"orientation"`
	AvgColor    string `json:"avgColor"`
}
