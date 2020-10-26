package test

// I know there are core libraries and external libraries to make it
// more compact, but I wanted to create the whole http.Client moching
// system to learn more and it was a good opportunity to do that.

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yitsushi/go-misskey"
)

// MockRequestHandler is just an alias for the function signature
// used to hanle mock requests.
type MockRequestHandler func(request *http.Request) (*http.Response, error)

// MockBody is a bulletproof ReadCloser.
type MockBody struct {
	*bytes.Reader
}

// Close the reader.
func (m MockBody) Close() error {
	return nil
}

// NewMockResponse creates a new mock response.
func NewMockResponse(code int, content []byte, err error) (*http.Response, error) {
	return &http.Response{
		StatusCode: code,
		Body: MockBody{
			bytes.NewReader(content),
		},
	}, err
}

// MockHTTPClient is a very basic mock implementation
// of the HTTPClient.
type MockHTTPClient struct {
	requests map[string]MockRequestHandler
}

// NewMockHTTPClient boostraps a new MockHTTPClient.
func NewMockHTTPClient() *MockHTTPClient {
	return &MockHTTPClient{
		requests: make(map[string]MockRequestHandler),
	}
}

// MockRequest registers a new endpoint.
func (c *MockHTTPClient) MockRequest(path string, handler MockRequestHandler) {
	c.requests[path] = handler
}

// Do is the real requirement to be an HTTPClient.
func (c *MockHTTPClient) Do(request *http.Request) (*http.Response, error) {
	if f, ok := c.requests[request.URL.Path]; ok {
		return f(request)
	}

	return NewMockResponse(http.StatusNotFound, []byte{}, nil)
}

// BadReadCloser is a basic ReadCloser with error. The only purpose
// is to test what is the connection dropped or something.
type BadReadCloser struct {
}

func (r BadReadCloser) Read(c []byte) (int, error) {
	return 0, errors.New("Read error") //nolint:goerr113
}

// Close the reader.
func (r BadReadCloser) Close() error {
	return errors.New("Close error") //nolint:goerr113
}

// SimpleMockOptions is the parameter list for SimpleMockEndpoint.
type SimpleMockOptions struct {
	Endpoint     string
	ResponseFile string
	RequestData  interface{}
	StatusCode   int
}

// SimpleMockEndpoint creates a simple MockHTTPClient that
// returns with an error if it was not able to pasre the request
// or returns with the content of the provided fixture file.
func SimpleMockEndpoint(options *SimpleMockOptions) *MockHTTPClient {
	mockClient := NewMockHTTPClient()
	mockClient.MockRequest(options.Endpoint, func(request *http.Request) (*http.Response, error) {
		defer request.Body.Close()
		body, _ := ioutil.ReadAll(request.Body)

		err := json.Unmarshal(body, options.RequestData)
		if err != nil {
			return NewMockResponse(
				http.StatusInternalServerError,
				[]byte(
					fmt.Sprintf(
						`{"error":{"message":"%s","code":"RANDOM_ERROR","kind":"client"}}`,
						err.Error(),
					),
				),
				err,
			)
		}

		return NewMockResponse(
			options.StatusCode,
			Must(LoadFixture(options.ResponseFile)),
			nil,
		)
	})

	return mockClient
}

// MakeMockClient creates a new Client with SimpleMockOptions.
func MakeMockClient(mockOptions SimpleMockOptions) *misskey.Client {
	mockClient := SimpleMockEndpoint(&mockOptions)
	client := misskey.NewClient("https://localhost", "thisistoken")
	client.HTTPClient = mockClient

	return client
}
