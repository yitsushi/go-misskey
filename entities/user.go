package entities

import "github.com/yitsushi/go-misskey/core"

// User is a... user.
type User struct {
	ID             string      `json:"id"`
	Name           string      `json:"name"`
	Username       string      `json:"username"`
	Host           core.String `json:"host"`
	AvatarURL      string      `json:"avatarUrl"`
	AvatarBlurhash core.String `json:"avatarBlurhash"`
	AvatarColor    core.String `json:"avatarColor"`
	Emojis         []Emoji     `json:"emojis"`
}
