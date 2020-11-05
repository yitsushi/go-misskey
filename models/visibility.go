package models

// Visibility is a simple string, but type alias for
// readability.
type Visibility string

const (
	// VisibilityPublic is for public notes.
	VisibilityPublic = "public"
	// VisibilityHome is for home notes.
	// It means, the post will be visible only on
	// your timeline.
	VisibilityHome = "home"
	// VisibilityFollowers is for only followers.
	VisibilityFollowers = "followers"
	// VisibilitySpecified is the case when other fields
	// defines who can see it.
	VisibilitySpecified = "specified"
)
