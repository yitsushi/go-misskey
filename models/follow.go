package models

// FollowStatus defines the status of followers and followees.
type FollowStatus struct {
	ID         string `json:"id"`
	CreatedAt  string `json:"createdAt"`
	FolloweeID string `json:"followeeId"`
	Followee   User   `json:"followee"`
	FollowerID string `json:"followerId"`
	Follower   User   `json:"follower,omitempty"`
}
