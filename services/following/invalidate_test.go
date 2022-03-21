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

func TestService_Invalidate(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/following/invalidate",
		RequestData:  &following.InvalidateRequest{},
		ResponseFile: "user.json",
		StatusCode:   http.StatusOK,
	})

	user, err := client.Following().Invalidate("88v9vu5nbu")
	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, "88v9vu5nbu", user.ID)
}

func TestInvalidateRequest_Validate(t *testing.T) {
	test.ValidateRequests(
		t,
		[]core.BaseRequest{
			following.InvalidateRequest{},
			following.InvalidateRequest{UserID: ""},
		},
		[]core.BaseRequest{
			following.InvalidateRequest{UserID: "88v9vu5nbu"},
		},
	)
}

func ExampleService_Invalidate() {
	client, _ := misskey.NewClientWithOptions(misskey.WithSimpleConfig("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN")))

	user, err := client.Following().Invalidate("88v9vu5nbu")
	if err != nil {
		log.Printf("[Following/Invalidate] %s", err)

		return
	}

	log.Printf("[Following/Invalidate] %s", user.Username)
}
