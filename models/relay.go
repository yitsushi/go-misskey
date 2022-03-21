package models

// RelayStatus is a type to represent relay statuses.
type RelayStatus string

const (
	// RelayStatusRequesting is the status of a pending relay.
	RelayStatusRequesting RelayStatus = "requesting"
	// RelayStatusAccepted is the status of an accepted relay.
	RelayStatusAccepted RelayStatus = "accepted"
	// RelayStatusRejected is the status of a rejected relay.
	RelayStatusRejected RelayStatus = "rejected"
)

// Relay is a simple representation of a Relay.
type Relay struct {
	ID     string      `json:"id"`
	Inbox  string      `json:"inbox"`
	Status RelayStatus `json:"status"`
}
