package queue

import (
	"github.com/yitsushi/go-misskey/core"
)

// Domain in the queue.
type Domain string

// State for a job.
type State string

// Job domains.
const (
	DeliverDomain       Domain = "deliver"
	InboxDomain         Domain = "inbox"
	DBDomain            Domain = "db"
	ObjectStorageDomain Domain = "objectStorage"
)

// Job states.
const (
	ActiveState  State = "active"
	WaitingState State = "waiting"
	DelayedState State = "delayed"
)

// Service is the base for all the endpoints on this service.
type Service struct {
	Call core.RequestHandlerFunc
}

// NewService creates a new Service instance.
func NewService(requestHandler core.RequestHandlerFunc) *Service {
	return &Service{Call: requestHandler}
}
