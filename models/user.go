package models

import (
	"time"

	"github.com/yitsushi/go-misskey/core"
)

// UserField is a custom field on user profiles.
type UserField struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// User is a... user.
type User struct {
	ID                             string      `json:"id"`
	Name                           string      `json:"name"`
	Username                       string      `json:"username"`
	Host                           core.String `json:"host"`
	AvatarURL                      string      `json:"avatarUrl"`
	AvatarBlurhash                 core.String `json:"avatarBlurhash"`
	AvatarColor                    core.String `json:"avatarColor"`
	CreatedAt                      *time.Time  `json:"createdAt"`
	UpdatedAt                      *time.Time  `json:"updatedAt"`
	Emojis                         []Emoji     `json:"emojis"`
	URL                            core.String `json:"url"`
	BannerURL                      string      `json:"BannerUrl"`
	BannerBlurhash                 core.String `json:"BannerBlurhash"`
	BannerColor                    core.String `json:"BannerColor"`
	IsAdmin                        bool        `json:"isAdmin"`
	IsModerator                    bool        `json:"isModerator"`
	IsBot                          bool        `json:"isBot"`
	IsCat                          bool        `json:"isCat"`
	IsLocked                       bool        `json:"isLocked"`
	IsSilenced                     bool        `json:"isSilenced"`
	IsSuspended                    bool        `json:"isSuspended"`
	TwoFactorEnabled               bool        `json:"twoFactorEnabled"`
	UsePasswordLessLogin           bool        `json:"usePasswordLessLogin"`
	SecurityKeys                   bool        `json:"securityKeys"`
	IsFollowing                    bool        `json:"isFollowing"`
	IsFollowed                     bool        `json:"isFollowed"`
	HasPendingFollowRequestFromYou bool        `json:"hasPendingFollowRequestFromYou"`
	HasPendingFollowRequestToYou   bool        `json:"hasPendingFollowRequestToYou"`
	IsBlocking                     bool        `json:"isBlocking"`
	IsBlocked                      bool        `json:"isBlocked"`
	IsMuted                        bool        `json:"isMuted"`
	Description                    string      `json:"description"`
	Location                       core.String `json:"location"`
	Birthday                       core.String `json:"birthday"`
	Fields                         []UserField `json:"fields"`
	FollowersCount                 uint64      `json:"followersCount"`
	FollowingCount                 uint64      `json:"followingCount"`
	NotesCount                     uint64      `json:"notesCount"`
	PinnedNoteIDs                  []string    `json:"pinnedNoteIds"`
	PinnedNotes                    []Note      `json:"pinnedNotes"`
	PinnedPageID                   core.String `json:"pinnedPageId"`
	PinnedPage                     core.String `json:"pinnedPage"`
}

// UserFromAdmin represents a user from admin/users/show.
// Usually user has an Emoji list of 'Emoji',
// but from /api/admin/show-user, it's a list of strings.
type UserFromAdmin struct {
	User

	InjectFeaturedNote bool     `json:"injectFeaturedNote"`
	AlwaysMarkNSFW     bool     `json:"alwaysMarkNsfw"`
	CarefulBot         bool     `json:"carefulBot"`
	AutoAcceptFollowed bool     `json:"autoAcceptFollowed"`
	EmailVerified      bool     `json:"emailVerified"`
	Email              string   `json:"email"`
	Token              string   `json:"token"`
	SecurityKeysList   []string `json:"securityKeysList"`

	Emojis []string `json:"emojis"`
}
