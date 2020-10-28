package hashtags

import "fmt"

// SortFlag is used for sorting on the List and Users endpoint.
// Here, there are two functions on it, Ascending and Descending.
// They specify how the tags or users should be sorted.
//
// In the background, the endpoint expects a string
// with a + or - sign.
type SortFlag string

const (
	// SortTagsByMentionedUsers sorts hashtags by user mentions.
	SortTagsByMentionedUsers SortFlag = "mentionedUsers"
	// SortTagsByMentionedLocalUsers sorts hashtags by local only user mentions.
	SortTagsByMentionedLocalUsers SortFlag = "mentionedLocalUsers"
	// SortTagsByMentionedRemoteUsers sorts hashtags by remote only user mentions.
	SortTagsByMentionedRemoteUsers SortFlag = "mentionedRemoteUsers"
	// SortTagsByAttachedUsers sorts hashtags by user attachment.
	SortTagsByAttachedUsers SortFlag = "attachedUsers"
	// SortTagsByAttachedLocalUsers sorts hashtags by local only user attachment.
	SortTagsByAttachedLocalUsers SortFlag = "attachedLocalUsers"
	// SortTagsByAttachedRemoteUsers sorts hashtags by remote only user attachment.
	SortTagsByAttachedRemoteUsers SortFlag = "attachedRemoteUsers"

	// SortUsersByFollowers sorts users by the number of their followers.
	SortUsersByFollowers SortFlag = "follower"
	// SortUsersByCreatedAt sorts users based on creation time.
	SortUsersByCreatedAt SortFlag = "createdAt"
	// SortUsersByUpdatedAt sorts users based on last update time.
	SortUsersByUpdatedAt SortFlag = "updatedAt"
)

// Descending order.
func (s SortFlag) Descending() string {
	return fmt.Sprintf("-%s", s)
}

// Ascending order.
func (s SortFlag) Ascending() string {
	return fmt.Sprintf("+%s", s)
}
