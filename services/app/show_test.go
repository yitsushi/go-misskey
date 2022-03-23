package app_test

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/services/app"
	"github.com/yitsushi/go-misskey/test"
)

func TestService_Show(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/app/show",
		RequestData:  &app.ShowRequest{},
		ResponseFile: "app.json",
		StatusCode:   http.StatusOK,
	})

	resp, err := client.App().Show("8y7eu4z91l")
	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, "test app", resp.Name)
}

func TestShowRequest_Validate(t *testing.T) {
	test.ValidateRequests(
		t,
		[]core.BaseRequest{
			app.ShowRequest{},
		},
		[]core.BaseRequest{
			app.ShowRequest{AppID: "asd"},
		},
	)
}

func ExampleService_Show() {
	client, _ := misskey.NewClientWithOptions(
		misskey.WithSimpleConfig("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN")),
	)

	resp, err := client.App().Show("8y7eu4z91l")
	if err != nil {
		log.Printf("[App/Show] %s", err)

		return
	}

	log.Printf("[App/Show] %s", resp.Name)
}
