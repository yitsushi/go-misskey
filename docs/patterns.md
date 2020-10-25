# Patterns

## Overview

Each endpoint group has its own Service under the `services` directory.  A
service may have sub-services for readability and better structure.

All global models are living under the `models` directory. There are only
structs, no functions (yet, maybe later). They represent a given resource like
User, Note or Page. Some of the endpoints are returning with a response where
some of the fields are not populated.

The `core` directory is a nice place for anything that has nothing to do with
endpoints, but they make the communication easier and helps services to make
calls without code duplication.

Note: The official documentation is just a guide, it has a lot of problems
      like most of the endpoints has no response body based on the
      documentation, but in reality, they have.

## Request Types

There are two types of requests, JSON and Multipart request. Most of the time
JSON request. The only exceptions are when we upload files, they are handled as
a `multipartform` request.

The client expects a struct that implements the `Request` interface. If we have
to add a new request type, it's enough to implement the `Request` interface.

There is a `BaseRequest` interface, all request handler expect structs with
`BaseRequest` implementation as they will call the `Validate` function on them.

For Multipart requests, we have a custom tag, so it's easy to write "forms" for
multipart form request.

Example multipart request struct:

```go
// CreateRequest represents a request to create a file.
type CreateRequest struct {
	FolderID    string `multipart:"folderId,type=field"`
	Name        string `multipart:"name,type=field"`
	IsSensitive bool   `multipart:"isSensitive,type=field"`
	Force       bool   `multipart:"force,type=field"`
	Content     []byte `multipart:"ref=name,type=file"`
}
```

The name of the field stands alone without `=`. If a field refers to another
one as their name (it's used for files, I don't think it has any other use), we
don't give it a name, but define `ref` and the value will be the referred name.
In this case, `Content` will be send as a file and the name of the file will be
the same as the value of the `Name` field.

## Service structure

Each service has to be registered in the `services.go` in the repository root.

Each service has a `service.go` file, that defines the service itself.  In this
file, we can define additional sub-services the same way as we register root
services. All global constants live there.

Example `service.go`:

```go
package something

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/services/something/stuff"
)

const (
	// DefaultListLimit is the default value for the limit lists.
	DefaultListLimit = 10
)

// Service is the base for all the endpoints on this service.
type Service struct {
	Call core.RequestHandlerFunc
}

// NewService creates a new Service instance.
func NewService(requestHandler core.RequestHandlerFunc) *Service {
	return &Service{Call: requestHandler}
}

// Stuff contains all endpoints under /something/stuff.
func (s *Service) Stuff() *files.Service {
	return files.NewService(s.Call)
}
```

Each request has its own file. Even if two requests are similar,
we don't reuse. All endpoints are exposed with their function
and their request.

Example endpoint:
```go
package something

import (
	"github.com/yitsushi/go-misskey/core"
)

// HelloRequest is doing something.
type HelloRequest struct{
    Name string `json:"id"`
}

// HelloResponse is the representation of the /drive/files request.
type HelloResponse struct {
	Message string `json:"message"`
}

// Validate the request.
func (r HelloRequest) Validate() error {
	if r.Name == "" {
		return core.RequestValidationError{
			Request: r,
			Message: core.UndefinedRequiredField,
			Field:   "Name",
		}
	}

	return nil
}

// Hello welcomes you.
func (s *Service) Hello(request HelloRequest) (HelloResponse, error) {
	var response HelloResponse
	err := s.Call(
		&core.JSONRequest{Request: &request, Path: "/something/hello"},
		&response,
	)

	return response, err
}
```

An Endpoint function can have a request argument or a single value, like on
delete, it's easier to use if we just ask for an ID, instead of a DeleteRequest
with only an ID field. In the function, we still create a DeleteRequest, but
the user of this library does not have to use that request as an argument for
the function.

## Testing

Right now, testing is kind of useless, but still has some value. Especially if
we want to extend or change something and we want to be sure everything is
working as intended. For the endpoint itself, there is a Mock server.  For
Validation, we can simply create a new Request and call the Validate function
and check for errors. The endpoint function test does not check for validation
errors, it will make a valid request always as all validation tests are tested
separately. For testing, we are using real responses from the server in fixture
files. They are living under the service directory (see any of the services,
there is a fixtures directory with files in it).

Each endpoint has a full example in their test file.

Example test:

```go
package something_test

import (
	"log"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/services/something"
	"github.com/yitsushi/go-misskey/test"
)

func TestService_Hello(t *testing.T) {
	mockClient := test.SimpleMockEndpoint(&test.SimpleMockOptions{
		Endpoint:     "/api/something/stuff",
		RequestData:  &something.HelloRequest{},
		ResponseFile: "create.json",
		StatusCode:   http.StatusOK,
	})

	client := misskey.NewClient("https://localhost", "thisistoken")
	client.HTTPClient = mockClient

	resp, err := client.Something().Hello(something.HelloRequest{
		Name: "xxxxxx"
	})
	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, "Hello xxxxxx", resp.Message)
}

func TestHelloRequest_Validate(t *testing.T) {
	testCase := something.HelloRequest{}
	assert.Error(t, testCase.Validate())
}

func ExampleService_Hello() {
	client := misskey.NewClient("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN"))

	resp, err := client.Something().Hello(something.HelloRequest{
		Name: "xxxxxx",
	})
	if err != nil {
		log.Printf("[Something/Hello] %s", err)

		return
	}

	log.Printf("[Something/Hello] %s", resp.Name)
}
```

## Get real world responses

As you write the request and the function itself, you can specify the response
as `core.DummyResponse`. It does nothing, it's an empty response struct and it
will fail on parsing the response JSON data, but with
`client.LogLevel(logrus.DebugLevel)`, you can see what was the response body.
Same it as a fixture file and you can write test for it and you can create an
accurate Response struct.

## Manual testing

You can create a simple program with a single `main` function where you can
call an endpoint.

```go
package main

import (
	"log"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/services/clips"
)

func main() {
	client := misskey.NewClient("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN"))
	client.LogLevel(logrus.DebugLevel)

	clips, err := client.Clips().Update(clips.UpdateRequest{
		ClipID: "8drxu3ckca",
		Name:   "new test",
	})
	if err != nil {
		log.Printf("[Clips] Error happened: %s", err)

		return
	}

	log.Println(clips)
}
```

This way, you can see (because of the LogLevel) the request body, the request
URL, the respnse body and response status code. It's an easy way to check if
everything is ok.
