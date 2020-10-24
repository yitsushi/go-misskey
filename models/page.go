package models

import "time"

// Page is the representation of a Page.
//
// Right now, as I'm working on Hashtags/Users, I don't know
// how Content, Variables and Summary are structured,
// so I just leave them as interface{} for now, and when
// we start to implement page endpoint, we can fill in the gap.
type Page struct {
	ID                  string     `json:"id"`
	CreatedAt           *time.Time `json:"createdAt"`
	UpdatedAt           *time.Time `json:"updatedAt"`
	UserID              string     `json:"userId"`
	User                *User      `json:"user"`
	Title               string     `json:"title"`
	Name                string     `json:"name"`
	HideTitleWhenPinned bool       `json:"hideTitleWhenPinned"`
	AlignCenter         bool       `json:"alignCenter"`
	IsLiked             bool       `json:"isLiked"`
	LikedCount          int        `json:"likedCount"`
	Font                string     `json:"font"`
	Script              string     `json:"script"`
	EyeCatchingImageID  string     `json:"eyeCatchingImageId"`
	EyeCatchingImage    *File      `json:"eyeCatchingImage"`
	AttachedFiles       []File     `json:"attachedFiles"`

	Content   interface{}   `json:"content"`
	Summary   interface{}   `json:"summary"`
	Variables []interface{} `json:"variables"`
}
