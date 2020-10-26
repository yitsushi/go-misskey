package models

import "time"

// Reaction is the representation of a Reaction.
type Reaction struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	User      User      `json:"user"`
	Type      string    `json:"type"`
}
