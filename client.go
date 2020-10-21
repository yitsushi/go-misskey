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
	"github.com/yitsushi/go-misskey/services/antennas"
	"github.com/yitsushi/go-misskey/services/meta"
	"golang.org/x/net/context"
)

// HTTPClient is a simple intreface for http.Client.
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// ClientInterface is an interface to describe how a Client looks like.
// Mostly for Mocking. Or later if Misskey gets multiple API versions.
type ClientInterface interface {
	Meta() *meta.Service
	Antennas() *antennas.Service
}

// Client is the main Misskey client struct.
type Client struct {
	BaseURL    string
	Token      string
	HTTPClient HTTPClient

	logger *logrus.Logger
}

// RequestTimout is the timeout of a request in seconds.
const RequestTimout = 10

// NewClient creates a new Misskey Client.
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

func (c *Client) LogLevel(level logrus.Level) {
	c.logger.SetLevel(level)
}

func (c Client) url(path string) string {
	return fmt.Sprintf("%s/api%s", c.BaseURL, path)
}

func (c Client) sendJSONRequest(request *core.BaseRequest, response interface{}) error {
	request.SetAPIToken(c.Token)

	requestBody, err := request.ToJSON()
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(
		context.Background(),
		"POST",
		c.url(request.Path),
		bytes.NewBuffer(requestBody),
	)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
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

	c.logger.WithField("_type", "response").WithField("from", req.URL).Debugf("%s", body)

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

	var errorWrapper core.ErrorResponseWrapper

	err = json.Unmarshal(body, &errorWrapper)
	if err != nil {
		return core.RequestError{Message: core.ErrorResponseParseError, Origin: err}
	}

	var errorResponse core.ErrorResponse
	if err := json.Unmarshal(errorWrapper.Error, &errorResponse); err != nil {
		return core.RequestError{Message: core.ErrorResponseParseError, Origin: err}
	}

	return core.UnknownError{Response: errorResponse}
}
