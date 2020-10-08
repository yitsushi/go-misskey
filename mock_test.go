package misskey_test

// I know there are core libraries and external libraries to make it
// more compact, but I wanted to create the whole http.Client moching
// system to learn more and it was a good opportunity to do that.

import (
	"bytes"
	"errors"
	"net/http"
)

type MockRequestHandler func(request *http.Request) (*http.Response, error)

type MockBody struct {
	*bytes.Reader
}

func (m MockBody) Close() error {
	return nil
}

func NewMockResponse(code int, content []byte, err error) (*http.Response, error) {
	return &http.Response{
		StatusCode: code,
		Body: MockBody{
			bytes.NewReader(content),
		},
	}, err
}

type MockHTTPClient struct {
	requests map[string]MockRequestHandler
}

func NewMockHTTPClient() *MockHTTPClient {
	return &MockHTTPClient{
		requests: make(map[string]MockRequestHandler),
	}
}

func (c *MockHTTPClient) MockRequest(path string, handler MockRequestHandler) {
	c.requests[path] = handler
}

func (c *MockHTTPClient) Do(request *http.Request) (*http.Response, error) {
	if f, ok := c.requests[request.URL.Path]; ok {
		return f(request)
	}

	return NewMockResponse(http.StatusNotFound, []byte{}, nil)
}

type BadReadCloser struct {
}

func (r BadReadCloser) Read(c []byte) (int, error) {
	return 0, errors.New("Read error") //nolint:goerr113
}

func (r BadReadCloser) Close() error {
	return errors.New("Close error") //nolint:goerr113
}
