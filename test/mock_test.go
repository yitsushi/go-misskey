package test_test

import (
	"bytes"
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/go-misskey/test"
)

type requestData struct {
	ID int `json:"id"`
}

func newClient() *test.MockHTTPClient {
	return test.SimpleMockEndpoint(&test.SimpleMockOptions{
		Endpoint:     "/something",
		ResponseFile: "data",
		RequestData:  &requestData{},
		StatusCode:   http.StatusOK,
	})
}

func TestSimpleMockEndpoint(t *testing.T) {
	client := newClient()

	req, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodPost,
		"/something",
		bytes.NewBufferString(`{"id": 12}`),
	)
	if !assert.NoError(t, err) {
		return
	}

	response, err := client.Do(req)
	if !assert.NoError(t, err) {
		return
	}

	defer response.Body.Close()
}

func TestSimpleMockEndpoint_invalidRequest(t *testing.T) {
	client := newClient()

	req, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodPost,
		"/something",
		bytes.NewBufferString(`booo`),
	)
	if !assert.NoError(t, err) {
		return
	}

	response, err := client.Do(req)
	if !assert.Error(t, err) {
		return
	}

	defer response.Body.Close()

	assert.Equal(t, http.StatusInternalServerError, response.StatusCode)
}

func TestSimpleMockEndpoint_noEndpoint(t *testing.T) {
	client := newClient()

	req, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodPost,
		"/something/else",
		bytes.NewBufferString(`{"id": 12}`),
	)
	if !assert.NoError(t, err) {
		return
	}

	response, err := client.Do(req)
	if !assert.NoError(t, err) {
		return
	}

	defer response.Body.Close()

	assert.Equal(t, http.StatusNotFound, response.StatusCode)
}
