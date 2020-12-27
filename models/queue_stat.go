package models

// QueueStat gives an overview of a queue with numbers.
type QueueStat struct {
	Waiting   int64 `json:"waiting"`
	Active    int64 `json:"active"`
	Completed int64 `json:"completed"`
	Failed    int64 `json:"failed"`
	Delayed   int64 `json:"delayed"`
	Paused    int64 `json:"paused"`
}

// QueueStats contains all available queue stats.
type QueueStats struct {
	Deliver       QueueStat `json:"deliver"`
	Inbox         QueueStat `json:"inbox"`
	DB            QueueStat `json:"db"`
	ObjectStorage QueueStat `json:"objectStorage"`
}
