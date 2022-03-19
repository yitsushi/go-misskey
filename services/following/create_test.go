package following_test

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/services/following"
	"github.com/yitsushi/go-misskey/test"
)

func TestService_Create(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/following/create",
		RequestData:  &following.CreateRequest{},
		ResponseFile: "user.json",
		StatusCode:   http.StatusOK,
	})

	user, err := client.Following().Create("88v9vu5nbu")
	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, "88v9vu5nbu", user.ID)
}

func TestCreateRequest_Validate(t *testing.T) {
	test.ValidateRequests(
		t,
		[]core.BaseRequest{
			following.CreateRequest{},
			following.CreateRequest{UserID: ""},
		},
		[]core.BaseRequest{
			following.CreateRequest{UserID: "88v9vu5nbu"},
		},
	)
}

func ExampleService_Create() {
	client := misskey.NewClient("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN"))

	user, err := client.Following().Create("88v9vu5nbu")
	if err != nil {
		log.Printf("[Following/Create] %s", err)

		return
	}

	log.Printf("[Following/Create] %s", user.Username)
}
