package users_test

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/services/admin/users"
	"github.com/yitsushi/go-misskey/test"
)

func TestService_Suspend(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/admin/suspend-user",
		RequestData:  &users.SuspendRequest{},
		ResponseFile: "empty",
		StatusCode:   http.StatusNoContent,
	})

	err := client.Admin().Users().Suspend("83sv4lyx22")
	assert.NoError(t, err)
}

func TestSuspendRequest_Validate(t *testing.T) {
	test.ValidateRequests(
		t,
		[]core.BaseRequest{
			users.SuspendRequest{},
		},
		[]core.BaseRequest{
			users.SuspendRequest{UserID: "asd"},
		},
	)
}

func ExampleService_Suspend() {
	client, _ := misskey.NewClientWithOptions(misskey.WithSimpleConfig("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN")))

	err := client.Admin().Users().Suspend("83sv4lyx22")
	if err != nil {
		log.Printf("[Admin/Users] %s", err)

		return
	}
}
