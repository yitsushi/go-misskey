package models

// Hashtag represents a hashtag.
type Hashtag struct {
	Tag                       string `json:"tag"`
	MentionedUsersCount       uint64 `json:"mentionedUsersCount"`
	MentionedLocalUsersCount  uint64 `json:"mentionedLocalUsersCount"`
	MentionedRemoteUsersCount uint64 `json:"mentionedRemoteUsersCount"`
	AttachedUsersCount        uint64 `json:"attachedUsersCount"`
	AttachedLocalUsersCount   uint64 `json:"attachedLocalUsersCount"`
	AttachedRemoteUsersCount  uint64 `json:"attachedRemoteUsersCount"`
}
