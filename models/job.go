package models

// Job is one job in the queue.
// Data is an interface for now, as it's not a unified resposne.
// The response is different based on the called Queue.
// For example deliver has a simple response, but inbox contains
// the Federated Activity object itself.
type Job struct {
	ID          string      `json:"id"`
	Attempts    int         `json:"attempts"`
	MaxAttempts int         `json:"maxAttempts"`
	Timestamp   int64       `json:"timestamp"`
	Data        interface{} `json:"data"`
}
