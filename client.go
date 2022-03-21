package misskey

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/yitsushi/go-misskey/core"
	"golang.org/x/net/context"
)

// Client is the main Misskey client struct.
type Client struct {
	BaseURL    string
	Token      string
	HTTPClient core.HTTPClient

	logger *logrus.Logger
}

// RequestTimout is the timeout of a request in seconds.
const RequestTimout = 10

// NewClient creates a new Misskey Client.
//
// Deprecated: use NewClientWithOptions instead.
func NewClient(baseURL, token string) *Client {
	return &Client{
		Token:   token,
		BaseURL: baseURL,
		HTTPClient: &http.Client{
			Timeout: time.Second * RequestTimout,
		},
		logger: logrus.New(),
	}
}

// NewClientWithOptions creates a new Misskey Client with defined options.
func NewClientWithOptions(options ...ClientOption) (*Client, error) {
	client := &Client{
		Token:      "",
		BaseURL:    "",
		HTTPClient: nil,
		logger:     logrus.New(),
	}

	for _, opt := range options {
		err := opt(client)
		if err != nil {
			return nil, err
		}
	}

	if client.HTTPClient == nil {
		client.HTTPClient = &http.Client{
			Timeout: time.Second * RequestTimout,
		}
	}

	return client, nil
}

// LogLevel sets logger level.
func (c *Client) LogLevel(level logrus.Level) {
	c.logger.SetLevel(level)
}

func (c Client) url(path string) string {
	return fmt.Sprintf("%s/api%s", c.BaseURL, path)
}

func (c Client) sendRequest(request core.Request, response interface{}) error {
	requestBody, contentType, err := request.ToBody(c.Token)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(
		context.Background(),
		"POST",
		c.url(request.EndpointPath()),
		bytes.NewBuffer(requestBody),
	)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", contentType)
	req.Header.Set("User-Agent", "Misskey Go SDK")
	c.logger.WithField("_type", "request").Debugf("%s %s", req.Method, req.URL)
	c.logger.WithField("_type", "request").Debugf("%s", requestBody)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return core.RequestError{Message: core.ResponseReadError, Origin: err}
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return core.RequestError{Message: core.ResponseReadBodyError, Origin: err}
	}

	c.logger.WithFields(logrus.Fields{
		"_type": "response",
		"from":  req.URL,
		"code":  resp.StatusCode,
	}).Debugf("%s", body)

	if resp.StatusCode == http.StatusOK {
		err = json.Unmarshal(body, response)

		return err
	}

	if resp.StatusCode == http.StatusNoContent {
		// Status code 204 considered as valid status code
		// if given operation was processed, no error occurred
		// but nothing to return, like delete resources.
		return nil
	}

	if resp.StatusCode == http.StatusNotFound {
		return core.EndpointNotFound{
			Endpoint: request.EndpointPath(),
		}
	}

	return unwrapError(body)
}

func unwrapError(body []byte) error {
	var errorWrapper core.ErrorResponseWrapper

	err := json.Unmarshal(body, &errorWrapper)
	if err != nil {
		return core.RequestError{Message: core.ErrorResponseParseError, Origin: err}
	}

	var errorResponse core.ErrorResponse
	if err := json.Unmarshal(errorWrapper.Error, &errorResponse); err != nil {
		return core.RequestError{Message: core.ErrorResponseParseError, Origin: err}
	}

	return core.UnknownError{Response: errorResponse}
}
