package meta_test

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/services/meta"
	"github.com/yitsushi/go-misskey/test"
)

func TestService_Announcements_auth(t *testing.T) {
	mockClient := test.SimpleMockEndpoint(&test.SimpleMockOptions{
		Endpoint:     "/api/announcements",
		RequestData:  &meta.AnnouncementsRequest{},
		ResponseFile: "auth/announcements.json",
		StatusCode:   http.StatusOK,
	})

	client := misskey.NewClient("https://localhost", "thisistoken")
	client.HTTPClient = mockClient
	// client.LogLevel(logrus.DebugLevel)

	response, err := client.Meta().Announcements(&meta.AnnouncementOptions{})
	if !assert.NoError(t, err) {
		return
	}

	if !assert.Len(t, response, 1) {
		return
	}

	assert.Equal(t, "8d44utwtj6", core.StringValue(response[0].ID))
	assert.NotNil(t, response[0].CreatedAt)
	assert.Nil(t, response[0].UpdatedAt)
	assert.Equal(t, "Sorry, if it disturbed you.", core.StringValue(response[0].Text))
	assert.Equal(t, "Test accouncement", core.StringValue(response[0].Title))
	assert.Nil(t, response[0].ImageURL)
	assert.True(t, response[0].IsRead)
}

func TestService_Announcements_anon(t *testing.T) {
	mockClient := test.SimpleMockEndpoint(&test.SimpleMockOptions{
		Endpoint:     "/api/announcements",
		RequestData:  &meta.AnnouncementsRequest{},
		ResponseFile: "anon/announcements.json",
		StatusCode:   http.StatusOK,
	})

	client := misskey.NewClient("https://localhost", "thisistoken")
	client.HTTPClient = mockClient

	response, err := client.Meta().Announcements(&meta.AnnouncementOptions{})
	if !assert.NoError(t, err) {
		return
	}

	if !assert.Len(t, response, 1) {
		return
	}

	assert.False(t, response[0].IsRead)
}

func ExampleService_Announcements() {
	client := misskey.NewClient("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN"))

	announcements, err := client.Meta().Announcements(
		&meta.AnnouncementOptions{
			WithUnreads: true,
			SinceID:     "",
			UntilID:     "",
		},
	)
	if err != nil {
		log.Printf("[Announcements] Error happened: %s", err)

		return
	}

	for _, announcement := range announcements {
		log.Printf("[Announcements] %s", core.StringValue(announcement.Title))
	}
}
