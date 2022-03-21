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

func TestService_Reject(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/following/requests/reject",
		RequestData:  &requests.RejectRequest{},
		ResponseFile: "empty",
		StatusCode:   http.StatusNoContent,
	})

	err := client.Following().Requests().Reject("88v9vu5nbu")
	if !assert.NoError(t, err) {
		return
	}
}

func TestRejectRequest_Validate(t *testing.T) {
	test.ValidateRequests(
		t,
		[]core.BaseRequest{
			requests.RejectRequest{},
			requests.RejectRequest{UserID: ""},
		},
		[]core.BaseRequest{
			requests.RejectRequest{UserID: "88v9vu5nbu"},
		},
	)
}

func ExampleService_Reject() {
	client, _ := misskey.NewClientWithOptions(misskey.WithSimpleConfig("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN")))

	err := client.Following().Requests().Reject("88v9vu5nbu")
	if err != nil {
		log.Printf("[Following/Requests/Reject] %s", err)

		return
	}

	log.Println("[Following/Requests/Reject] Rejected")
}
