package models

// Following is a single following record.
// Technically it's the same as Followers but we keep it separate
// for easy tracking.
type Following struct {
	Followers
}

// Followers is a single follower record.
type Followers struct {
	ID         string `json:"id"`
	CreatedAt  string `json:"createdAt"`
	FolloweeID string `json:"followeeId"`
	Followee   User   `json:"followee"`
	FollowerID string `json:"followerId"`
	Follower   User   `json:"follower,omitempty"`
}
