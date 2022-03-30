package models

import "time"

// Poll represents a Poll data structure for Misskey.
type Poll struct {
	Multiple  bool      `json:"multiple"`
	ExpiresAt time.Time `json:"expiresAt"`
	Choices   []Choice  `json:"choices"`
}

// Choice is a sinple Choice in a Poll.
type Choice struct {
	Text    string `json:"text"`
	Votes   uint64 `json:"votes"`
	IsVoted bool   `json:"isVoted"`
}
