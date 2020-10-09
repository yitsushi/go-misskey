package misskey

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/services/antennas"
	"github.com/yitsushi/go-misskey/services/meta"
)

func (c *Client) requestHandler(request *core.BaseRequest, response interface{}) error {
	err := c.sendJSONRequest(
		request,
		response,
	)

	return err
}

func (c *Client) Meta() *meta.Service {
	return meta.NewService(c.requestHandler)
}

func (c *Client) Antennas() *antennas.Service {
	return antennas.NewService(c.requestHandler)
}
