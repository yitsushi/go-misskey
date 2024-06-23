package models

import "time"

// Invite has information about a single invite.
type Invite struct {
	ID        string    `json:"id"`
	Code      string    `json:"code"`
	ExpiredAt time.Time `json:"expiresAt"`
	CreatedAt time.Time `json:"createdAt"`
	CreatedBy *User     `json:"createdBy"`
	UsedBy    *User     `json:"usedBy"`
	UsedAt    time.Time `json:"usedAt"`
	Used      bool      `json:"used"`
}
