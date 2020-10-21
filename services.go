package misskey

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/services/antennas"
	"github.com/yitsushi/go-misskey/services/federation"
	"github.com/yitsushi/go-misskey/services/meta"
)

func (c *Client) requestHandler(request *core.BaseRequest, response interface{}) error {
	err := c.sendJSONRequest(
		request,
		response,
	)

	return err
}

// Meta is all the endpoints under Meta in the documentation.
// They don't have an API pth prefix.
func (c *Client) Meta() *meta.Service {
	return meta.NewService(c.requestHandler)
}

// Antennas contains all endpoints under /antennas.
func (c *Client) Antennas() *antennas.Service {
	return antennas.NewService(c.requestHandler)
}

// Federation contains all endpoints under /federation.
func (c *Client) Federation() *federation.Service {
	return federation.NewService(c.requestHandler)
}
