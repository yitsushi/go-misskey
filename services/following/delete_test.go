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

func TestService_Delete(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/following/delete",
		RequestData:  &following.DeleteRequest{},
		ResponseFile: "user.json",
		StatusCode:   http.StatusOK,
	})

	user, err := client.Following().Delete("88v9vu5nbu")
	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, "88v9vu5nbu", user.ID)
}

func TestDeleteRequest_Validate(t *testing.T) {
	test.ValidateRequests(
		t,
		[]core.BaseRequest{
			following.DeleteRequest{},
			following.DeleteRequest{UserID: ""},
		},
		[]core.BaseRequest{
			following.DeleteRequest{UserID: "88v9vu5nbu"},
		},
	)
}

func ExampleService_Delete() {
	client := misskey.NewClient("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN"))

	user, err := client.Following().Delete("88v9vu5nbu")
	if err != nil {
		log.Printf("[Following/Delete] %s", err)

		return
	}

	log.Printf("[Following/Delete] %s", user.Username)
}
