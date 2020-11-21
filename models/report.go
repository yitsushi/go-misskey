package models

import "time"

// Report represents a reported user.
// Only users can be reported, for reports on specific
// resource like Note, the Comment field container
// the URL for that resource.
type Report struct {
	ID           string    `json:"id"`
	CreatedAt    time.Time `json:"createdAt"`
	Comment      string    `json:"comment"`
	Resolved     bool      `json:"resolved"`
	ReporterID   string    `json:"reporterId"`
	TargetUserID string    `json:"targetUserId"`
	AssigneeID   string    `json:"assigneeId"`
	Reporter     *User     `json:"reporter"`
	TargetUser   *User     `json:"targetUser"`
	Assignee     *User     `json:"assignee"`
}
