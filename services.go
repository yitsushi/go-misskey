package misskey

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/services/admin"
	"github.com/yitsushi/go-misskey/services/antennas"
	"github.com/yitsushi/go-misskey/services/clips"
	"github.com/yitsushi/go-misskey/services/drive"
	"github.com/yitsushi/go-misskey/services/federation"
	"github.com/yitsushi/go-misskey/services/hashtags"
	"github.com/yitsushi/go-misskey/services/meta"
	"github.com/yitsushi/go-misskey/services/notes"
	"github.com/yitsushi/go-misskey/services/notifications"
	"github.com/yitsushi/go-misskey/services/promo"
)

func (c *Client) requestHandler(request core.Request, response interface{}) error {
	if err := request.Validate(); err != nil {
		return err
	}

	return c.sendRequest(request, response)
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

// Notifications contains all endpoints under /notifications.
func (c *Client) Notifications() *notifications.Service {
	return notifications.NewService(c.requestHandler)
}

// Hashtags contains all endpoints under /hashtags.
func (c *Client) Hashtags() *hashtags.Service {
	return hashtags.NewService(c.requestHandler)
}

// Clips contains all endpoints under /clips.
func (c *Client) Clips() *clips.Service {
	return clips.NewService(c.requestHandler)
}

// Drive contains all endpoints under /drive.
func (c *Client) Drive() *drive.Service {
	return drive.NewService(c.requestHandler)
}

// Federation contains all endpoints under /federation.
func (c *Client) Federation() *federation.Service {
	return federation.NewService(c.requestHandler)
}

// Notes contains all endpoints under /notes.
func (c *Client) Notes() *notes.Service {
	return notes.NewService(c.requestHandler)
}

// Promo contains all endpoints under /promo.
func (c *Client) Promo() *promo.Service {
	return promo.NewService(c.requestHandler)
}

// Admin contains all endpoints under /admin.
func (c *Client) Admin() *admin.Service {
	return admin.NewService(c.requestHandler)
}
