package notifications_test

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/services/notifications"
	"github.com/yitsushi/go-misskey/test"
)

func TestService_MarkAllAsRead(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/notifications/mark-all-as-read",
		RequestData:  &notifications.MarkAllAsReadRequest{},
		ResponseFile: "empty",
		StatusCode:   http.StatusNoContent,
	})

	err := client.Notifications().MarkAllAsRead()
	if !assert.NoError(t, err) {
		return
	}
}

func ExampleService_MarkAllAsRead() {
	client, _ := misskey.NewClientWithOptions(misskey.WithSimpleConfig("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN")))

	err := client.Notifications().MarkAllAsRead()
	if err != nil {
		log.Printf("[Notifications] Error happened: %s", err)

		return
	}
}
