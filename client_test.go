package misskey_test

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/yitsushi/go-misskey"
)

func TestNewClient_NormalRequestContent(t *testing.T) {
	mockClient := NewMockHTTPClient()
	mockClient.MockRequest("/api/stats", func(request *http.Request) (*http.Response, error) {
		defer request.Body.Close()
		body, _ := ioutil.ReadAll(request.Body)

		var statsRequest map[string]interface{}

		err := json.Unmarshal(body, &statsRequest)
		if err != nil {
			t.Errorf("Unable to parse request: %s", err)
			return NewMockResponse(http.StatusInternalServerError, []byte{}, err)
		}

		if statsRequest["i"] != "thisistoken" {
			t.Errorf("expected api token = thisistoken; got = %s", statsRequest["i"])
		}

		return NewMockResponse(http.StatusOK, []byte("{}"), nil)
	})

	client := misskey.NewClient("https://localhost", "thisistoken")
	client.HTTPClient = mockClient

	_, err := client.Stats()
	if err != nil {
		t.Errorf("Unexpected error = %s", err)
	}
}

func TestNewClient_RequestError(t *testing.T) {
	mockClient := NewMockHTTPClient()
	mockClient.MockRequest("/api/stats", func(request *http.Request) (*http.Response, error) {
		return NewMockResponse(http.StatusNotImplemented, []byte{}, errors.New("bad")) //nolint:goerr113
	})

	client := misskey.NewClient("https://localhost", "thisistoken")
	client.HTTPClient = mockClient

	_, err := client.Stats()
	if err == nil {
		t.Error("Expected error, but never happened")
		return
	}

	expected := misskey.RequestError{
		Message: misskey.ResponseReadError,
		Origin:  errors.New("bad"), //nolint:goerr113
	}
	if err.Error() != expected.Error() {
		t.Errorf("Expected error = %s, got = %s", expected.Error(), err.Error())
	}
}

func TestNewClient_ReadError(t *testing.T) {
	mockClient := NewMockHTTPClient()
	mockClient.MockRequest("/api/stats", func(request *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       BadReadCloser{},
		}, nil
	})

	client := misskey.NewClient("https://localhost", "thisistoken")
	client.HTTPClient = mockClient

	_, err := client.Stats()
	if err == nil {
		t.Error("Expected error, but never happened")
		return
	}

	expected := misskey.RequestError{
		Message: misskey.ResponseReadBodyError,
		Origin:  errors.New("Read error"), //nolint:goerr113
	}
	if err.Error() != expected.Error() {
		t.Errorf("Expected error = %s, got = %s", expected.Error(), err.Error())
	}
}

func TestNewClient_ErrorResponseWrapper_Error(t *testing.T) {
	mockClient := NewMockHTTPClient()
	mockClient.MockRequest("/api/stats", func(request *http.Request) (*http.Response, error) {
		content := []byte("something")
		return NewMockResponse(http.StatusInternalServerError, content, nil)
	})

	client := misskey.NewClient("https://localhost", "thisistoken")
	client.HTTPClient = mockClient

	_, err := client.Stats()
	if err == nil {
		t.Error("Expected error, but never happened")
		return
	}

	expected := misskey.RequestError{
		Message: misskey.ErrorResponseParseError,
		Origin:  errors.New("invalid character 's' looking for beginning of value"), //nolint:goerr113
	}
	if err.Error() != expected.Error() {
		t.Errorf("Expected error = %s, got = %s", expected.Error(), err.Error())
	}
}

func TestNewClient_ErrorResponseParse_Error(t *testing.T) {
	mockClient := NewMockHTTPClient()
	mockClient.MockRequest("/api/stats", func(request *http.Request) (*http.Response, error) {
		content := []byte("{\"error\": true}")
		return NewMockResponse(http.StatusInternalServerError, content, nil)
	})

	client := misskey.NewClient("https://localhost", "thisistoken")
	client.HTTPClient = mockClient

	_, err := client.Stats()
	if err == nil {
		t.Error("Expected error, but never happened")
		return
	}

	expected := misskey.RequestError{
		Message: misskey.ErrorResponseParseError,
		Origin:  errors.New("json: cannot unmarshal bool into Go value of type misskey.ErrorResponse"), //nolint:goerr113
	}
	if err.Error() != expected.Error() {
		t.Errorf("Expected error = %s, got = %s", expected.Error(), err.Error())
	}
}

func TestNewClient_ValidErrorResponse(t *testing.T) {
	mockClient := NewMockHTTPClient()
	mockClient.MockRequest("/api/stats", func(request *http.Request) (*http.Response, error) {
		content := []byte(`{
			"error": {
				"info": {
					"param": "field",
					"reason": "this is the reason"
				}
			}
		}`)
		return NewMockResponse(http.StatusInternalServerError, content, nil)
	})

	client := misskey.NewClient("https://localhost", "thisistoken")
	client.HTTPClient = mockClient

	_, err := client.Stats()
	if err == nil {
		t.Error("Expected error, but never happened")
		return
	}

	expectedResponse := misskey.ErrorResponse{}
	expectedResponse.Info.Param = "field"
	expectedResponse.Info.Reason = "this is the reason"

	expected := misskey.UnknownError{
		Response: expectedResponse,
	}
	if err.Error() != expected.Error() {
		t.Errorf("Expected error = %s, got = %s", expected.Error(), err.Error())
	}
}
