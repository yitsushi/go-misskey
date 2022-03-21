package misskey

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

// ClientOption is a function that can be used to configure a client.
type ClientOption func(*Client) error

// ClientOptionError occures when something goes wrong with any of the requested
// options.
type ClientOptionError struct {
	Message string
}

func (e ClientOptionError) Error() string {
	return fmt.Sprintf("client options error: %s", e.Message)
}

// WithAPIToken configures the API token on the client.
func WithAPIToken(token string) ClientOption {
	return func(client *Client) error {
		client.Token = token

		return nil
	}
}

// WithBaseURL configures the base url of the Misskey instance.
//
// - Protocol: http, https
// - Domain: Well, that's the domain name
// - Path: Leave it empty, unless the target instance is not served from the
//         root path. Important: Do not add a tailing slash.
func WithBaseURL(protocol, domain, path string) ClientOption {
	return func(client *Client) error {
		if domain == "" {
			return ClientOptionError{Message: "undefined value: domain"}
		}

		if protocol == "" {
			protocol = "https"
		}

		if path != "" && path[0] != '/' {
			path = "/" + path
		}

		client.BaseURL = fmt.Sprintf("%s://%s%s", protocol, domain, path)

		return nil
	}
}

// WithLogLevel configures the logger to use the specified log level.
func WithLogLevel(level logrus.Level) ClientOption {
	return func(client *Client) error {
		client.LogLevel(level)

		return nil
	}
}

// WithSimpleConfig configures the client with similar logic as NewClient().
//
// The sole purpose of this to make it easier to migrate to the new function.
func WithSimpleConfig(baseURL, token string) ClientOption {
	return func(client *Client) error {
		client.BaseURL = baseURL
		client.Token = token

		return nil
	}
}
