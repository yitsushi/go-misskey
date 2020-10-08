package misskey

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/meta"
	"golang.org/x/net/context"
)

// HTTPClient is a simple intreface for http.Client.
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// ClientInterface is an interface to describe how a Client looks like.
// Mostly for Mocking. Or later if Misskey gets multiple API versions.
type ClientInterface interface {
	Announcements(options *AnnouncementOptions) (meta.AnnouncementsResponse, error)
	Meta(details bool) (meta.InstanceMetaResponse, error)
	Stats() (meta.StatsResponse, error)
}

// Client is the main Misskey client struct.
type Client struct {
	BaseURL    string
	Token      string
	HTTPClient HTTPClient
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
	}
}

func (c Client) url(path string) string {
	return fmt.Sprintf("%s/api%s", c.BaseURL, path)
}

func (c Client) sendJSONRequest(request *core.BaseRequest, respose interface{}) error {
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

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return RequestError{Message: ResponseReadError, Origin: err}
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return RequestError{Message: ResponseReadBodyError, Origin: err}
	}

	if resp.StatusCode == http.StatusOK {
		err = json.Unmarshal(body, respose)
		return err
	}

	var errorWrapper errorResponseWrapper

	err = json.Unmarshal(body, &errorWrapper)
	if err != nil {
		return RequestError{Message: ErrorResponseParseError, Origin: err}
	}

	var errorResponse ErrorResponse
	if err := json.Unmarshal(errorWrapper.Error, &errorResponse); err != nil {
		return RequestError{Message: ErrorResponseParseError, Origin: err}
	}

	return UnknownError{Response: errorResponse}
}
