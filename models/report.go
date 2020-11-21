package models

import "time"

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
