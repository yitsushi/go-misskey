package core

// BaseRequest is the base request.
type BaseRequest interface {
	Validate() error
}
