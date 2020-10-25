package models

import "time"

// Clip is a simple representation of a Clip.
type Clip struct {
	ID        string     `json:"id"`
	CreatedAt *time.Time `json:"createdAt"`
	Name      string     `json:"name"`
}
