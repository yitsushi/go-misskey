package moderation

import (
	"github.com/yitsushi/go-misskey/core"
)

// ReportState is the state of a report.
type ReportState string

// UserOrigin of a user.
type UserOrigin string

const (
	// AllState for all reports.
	AllState ReportState = ""
	// ResolvedState for resolved reports.
	ResolvedState ReportState = "resolved"
	// UnresolvedState for unresolved reports.
	UnresolvedState ReportState = "unresolved"

	// OriginCombined for local and remote users.
	OriginCombined UserOrigin = "combined"
	// OriginLocal for only local users.
	OriginLocal UserOrigin = "local"
	// OriginRemote for only remote users.
	OriginRemote UserOrigin = "remote"
)

// Service is the base for all the endpoints on this service.
type Service struct {
	Call core.RequestHandlerFunc
}

// NewService creates a new Service instance.
func NewService(requestHandler core.RequestHandlerFunc) *Service {
	return &Service{Call: requestHandler}
}
