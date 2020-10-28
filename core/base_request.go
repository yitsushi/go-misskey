package core

// BaseRequest is used to add simple functionality to all requests. Each request
// is responsible for its own validation by implementing this interface.
type BaseRequest interface {
	Validate() error
}
