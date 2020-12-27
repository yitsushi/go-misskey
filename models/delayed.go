package models

// Delayed job count per host.
type Delayed struct {
	Host  string
	Count int64
}
