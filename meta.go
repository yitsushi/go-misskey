package misskey

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/meta"
)

// AnnouncementOptions is the options list for Announcement().
type AnnouncementOptions struct {
	WithUnreads bool
	SinceID     string
	UntilID     string
}

// Announcements lists all announcements.
func (c *Client) Announcements(options *AnnouncementOptions) (meta.AnnouncementsResponse, error) {
	request := &meta.AnnouncementsRequest{
		WithUnreads: options.WithUnreads,
		SinceID:     options.SinceID,
		UntilID:     options.UntilID,
	}

	var respose meta.AnnouncementsResponse

	err := c.sendJSONRequest(
		&core.BaseRequest{Request: request, Path: "/announcements"},
		&respose,
	)

	return respose, err
}

func (c *Client) Meta(details bool) (meta.InstanceMetaResponse, error) {
	request := &meta.InstanceMetaRequest{
		Detail: details,
	}

	var respose meta.InstanceMetaResponse

	err := c.sendJSONRequest(
		&core.BaseRequest{Request: request, Path: "/meta"},
		&respose,
	)

	return respose, err
}

func (c *Client) Stats() (meta.StatsResponse, error) {
	var response meta.StatsResponse

	err := c.sendJSONRequest(
		&core.BaseRequest{Request: &meta.StatsRequest{}, Path: "/stats"},
		&response,
	)

	return response, err
}
