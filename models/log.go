package models

import "time"

// Log from admin/logs.
type Log struct {
	ID        string      `json:"id"`
	CreatedAt time.Time   `json:"createdAt"`
	Domain    []string    `json:"domain"`
	Level     string      `json:"level"`
	Worker    string      `json:"worker"`
	Machine   string      `json:"machine"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
}

// ModerationLog from admin/show-moderation-logs.
type ModerationLog struct {
	ID        string      `json:"id"`
	CreatedAt time.Time   `json:"createdAt"`
	Type      string      `json:"type"`
	UserID    string      `json:"userId"`
	User      User        `json:"user"`
	Info      interface{} `json:"info"`
}
