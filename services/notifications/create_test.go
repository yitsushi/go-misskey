package notifications_test

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/services/notifications"
	"github.com/yitsushi/go-misskey/test"
)

func TestService_Create(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/notifications/create",
		RequestData:  &notifications.CreateRequest{},
		ResponseFile: "empty",
		StatusCode:   http.StatusNoContent,
	})

	err := client.Notifications().Create(notifications.CreateRequest{
		Body:   "This is the body",
		Header: core.NewString("This is the header"),
	})
	if !assert.NoError(t, err) {
		return
	}
}

func ExampleService_Create() {
	client := misskey.NewClient("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN"))

	err := client.Notifications().Create(notifications.CreateRequest{
		Header: core.NewString("Thi is the header"),
		Body:   "Example notification",
	})
	if err != nil {
		log.Printf("[Notifications] Error happened: %s", err)

		return
	}
}
