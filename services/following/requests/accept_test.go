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

func TestService_Accept(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/following/requests/accept",
		RequestData:  &requests.AcceptRequest{},
		ResponseFile: "empty",
		StatusCode:   http.StatusNoContent,
	})

	err := client.Following().Requests().Accept("88v9vu5nbu")
	if !assert.NoError(t, err) {
		return
	}
}

func TestAcceptRequest_Validate(t *testing.T) {
	test.ValidateRequests(
		t,
		[]core.BaseRequest{
			requests.AcceptRequest{},
			requests.AcceptRequest{UserID: ""},
		},
		[]core.BaseRequest{
			requests.AcceptRequest{UserID: "88v9vu5nbu"},
		},
	)
}

func ExampleService_Accept() {
	client, _ := misskey.NewClientWithOptions(misskey.WithSimpleConfig("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN")))

	err := client.Following().Requests().Accept("88v9vu5nbu")
	if err != nil {
		log.Printf("[Following/Requests/Accept] %s", err)

		return
	}

	log.Println("[Following/Requests/Accept] Accepted")
}
