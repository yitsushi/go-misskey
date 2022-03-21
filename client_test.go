package misskey_test

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/test"
)

func TestNewClient_NormalRequestContent(t *testing.T) {
	mockClient := test.NewMockHTTPClient()
	mockClient.MockRequest("/api/stats", func(request *http.Request) (*http.Response, error) {
		defer request.Body.Close()
		body, _ := ioutil.ReadAll(request.Body)

		var statsRequest map[string]interface{}

		err := json.Unmarshal(body, &statsRequest)
		if err != nil {
			t.Errorf("Unable to parse request: %s", err)

			return test.NewMockResponse(http.StatusInternalServerError, []byte{}, err)
		}

		if statsRequest["i"] != "thisistoken" {
			t.Errorf("expected api token = thisistoken; got = %s", statsRequest["i"])
		}

		return test.NewMockResponse(http.StatusOK, []byte("{}"), nil)
	})

	client, _ := misskey.NewClientWithOptions(
		misskey.WithAPIToken("thisistoken"),
		misskey.WithBaseURL("https", "localhost", ""),
		misskey.WithHTTPClient(mockClient),
	)

	_, err := client.Meta().Stats()
	if err != nil {
		t.Errorf("Unexpected error = %s", err)
	}
}

func TestNewClient_undefinedDomain(t *testing.T) {
	client, err := misskey.NewClientWithOptions(
		misskey.WithLogLevel(logrus.DebugLevel),
		misskey.WithBaseURL("", "", ""),
	)

	if err == nil {
		t.Error("Expected error, but never happened")

		return
	}

	expectedErrorMessage := "client options error: undefined value: domain"
	if err.Error() != expectedErrorMessage {
		t.Errorf("expected error = %s; got = %s", expectedErrorMessage, err.Error())
	}

	if client != nil {
		t.Error("NewClientWithOptions should return nil as client if error happened")
	}
}

func TestNewClient_RequestError(t *testing.T) {
	mockClient := test.NewMockHTTPClient()
	mockClient.MockRequest("/api/stats", func(request *http.Request) (*http.Response, error) {
		return test.NewMockResponse(http.StatusNotImplemented, []byte{}, errors.New("bad")) //nolint:goerr113
	})

	client, _ := misskey.NewClientWithOptions(
		misskey.WithSimpleConfig("https://localhost", "thisistoken"),
		misskey.WithHTTPClient(mockClient),
	)

	_, err := client.Meta().Stats()
	if err == nil {
		t.Error("Expected error, but never happened")

		return
	}

	expected := core.RequestError{
		Message: core.ResponseReadError,
		Origin:  errors.New("bad"), //nolint:goerr113
	}
	if err.Error() != expected.Error() {
		t.Errorf("Expected error = %s, got = %s", expected.Error(), err.Error())
	}
}

func TestNewClient_ReadError(t *testing.T) {
	mockClient := test.NewMockHTTPClient()
	mockClient.MockRequest("/api/stats", func(request *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       test.BadReadCloser{},
		}, nil
	})

	client, _ := misskey.NewClientWithOptions(
		misskey.WithAPIToken("thisistoken"),
		misskey.WithBaseURL("https", "localhost", ""),
		misskey.WithHTTPClient(mockClient),
	)

	_, err := client.Meta().Stats()
	if err == nil {
		t.Error("Expected error, but never happened")

		return
	}

	expected := core.RequestError{
		Message: core.ResponseReadBodyError,
		Origin:  errors.New("Read error"), //nolint:goerr113
	}
	if err.Error() != expected.Error() {
		t.Errorf("Expected error = %s, got = %s", expected.Error(), err.Error())
	}
}

func TestNewClient_ErrorResponseWrapper_Error(t *testing.T) {
	mockClient := test.NewMockHTTPClient()
	mockClient.MockRequest("/api/stats", func(request *http.Request) (*http.Response, error) {
		content := []byte("something")

		return test.NewMockResponse(http.StatusInternalServerError, content, nil)
	})

	client, _ := misskey.NewClientWithOptions(
		misskey.WithAPIToken("thisistoken"),
		misskey.WithBaseURL("https", "localhost", ""),
		misskey.WithHTTPClient(mockClient),
	)

	_, err := client.Meta().Stats()
	if err == nil {
		t.Error("Expected error, but never happened")

		return
	}

	expected := core.RequestError{
		Message: core.ErrorResponseParseError,
		Origin:  errors.New("invalid character 's' looking for beginning of value"), //nolint:goerr113
	}
	if err.Error() != expected.Error() {
		t.Errorf("Expected error = %s, got = %s", expected.Error(), err.Error())
	}
}

func TestNewClient_ErrorResponseParse_Error(t *testing.T) {
	mockClient := test.NewMockHTTPClient()
	mockClient.MockRequest("/api/stats", func(request *http.Request) (*http.Response, error) {
		content := []byte("{\"error\": true}")

		return test.NewMockResponse(http.StatusInternalServerError, content, nil)
	})

	client, _ := misskey.NewClientWithOptions(
		misskey.WithAPIToken("thisistoken"),
		misskey.WithBaseURL("https", "localhost", ""),
		misskey.WithHTTPClient(mockClient),
	)

	_, err := client.Meta().Stats()
	if err == nil {
		t.Error("Expected error, but never happened")

		return
	}

	expected := core.RequestError{
		Message: core.ErrorResponseParseError,
		Origin:  errors.New("json: cannot unmarshal bool into Go value of type core.ErrorResponse"), //nolint:goerr113
	}
	if err.Error() != expected.Error() {
		t.Errorf("Expected error = %s, got = %s", expected.Error(), err.Error())
	}
}

func TestNewClient_ValidErrorResponse(t *testing.T) {
	mockClient := test.NewMockHTTPClient()
	mockClient.MockRequest("/api/stats", func(request *http.Request) (*http.Response, error) {
		content := []byte(`{
			"error": {
				"info": {
					"param": "field",
					"reason": "this is the reason"
				}
			}
		}`)

		return test.NewMockResponse(http.StatusInternalServerError, content, nil)
	})

	client, _ := misskey.NewClientWithOptions(
		misskey.WithAPIToken("thisistoken"),
		misskey.WithBaseURL("https", "localhost", ""),
		misskey.WithHTTPClient(mockClient),
	)

	_, err := client.Meta().Stats()

	if err == nil {
		t.Error("Expected error, but never happened")

		return
	}

	expectedResponse := core.ErrorResponse{}
	expectedResponse.Info.Param = "field"
	expectedResponse.Info.Reason = "this is the reason"

	expected := core.UnknownError{
		Response: expectedResponse,
	}
	if err.Error() != expected.Error() {
		t.Errorf("Expected error = %s, got = %s", expected.Error(), err.Error())
	}
}
