package models

import "time"

// Group represents a user group.
type Group struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`
	OwnerID   string    `json:"ownerId"`
	UserIDs   []string  `json:"userIds"`
}
