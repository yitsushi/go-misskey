package core

import "net/http"

// HTTPClient is a simple intreface for http.Client.
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
