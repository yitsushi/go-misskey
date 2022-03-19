package requests_test

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/services/following/requests"
	"github.com/yitsushi/go-misskey/test"
)

func TestService_Cancel(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/following/requests/cancel",
		RequestData:  &requests.CancelRequest{},
		ResponseFile: "user.json",
		StatusCode:   http.StatusOK,
	})

	user, err := client.Following().Requests().Cancel("88v9vu5nbu")
	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, "88v9vu5nbu", user.ID)
}

func TestCancelRequest_Validate(t *testing.T) {
	test.ValidateRequests(
		t,
		[]core.BaseRequest{
			requests.CancelRequest{},
			requests.CancelRequest{UserID: ""},
		},
		[]core.BaseRequest{
			requests.CancelRequest{UserID: "88v9vu5nbu"},
		},
	)
}

func ExampleService_Cancel() {
	client := misskey.NewClient("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN"))

	user, err := client.Following().Requests().Cancel("88v9vu5nbu")
	if err != nil {
		log.Printf("[Following/Requests/Cancel] %s", err)

		return
	}

	log.Printf("[Following/Requests/Cancel] Canceled: %s", user.ID)
}
