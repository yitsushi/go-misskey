package core

// Request is the interface that has to be implemented to be valid request.
type Request interface {
	ToBody(token string) (body []byte, contentType string, err error)
	EndpointPath() string
}
