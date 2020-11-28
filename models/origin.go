package models

// UserOrigin of a user.
type UserOrigin string

const (
	// OriginCombined for local and remote users.
	OriginCombined UserOrigin = "combined"
	// OriginLocal for only local users.
	OriginLocal UserOrigin = "local"
	// OriginRemote for only remote users.
	OriginRemote UserOrigin = "remote"
)
