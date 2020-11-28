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

func TestService_Unsuspend(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/admin/unsuspend-user",
		RequestData:  &users.UnsuspendRequest{},
		ResponseFile: "empty",
		StatusCode:   http.StatusNoContent,
	})

	err := client.Admin().Users().Unsuspend("83sv4lyx22")
	assert.NoError(t, err)
}

func TestUnsuspendRequest_Validate(t *testing.T) {
	test.ValidateRequests(
		t,
		[]core.BaseRequest{
			users.UnsuspendRequest{},
		},
		[]core.BaseRequest{
			users.UnsuspendRequest{UserID: "asd"},
		},
	)
}

func ExampleService_Unsuspend() {
	client := misskey.NewClient("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN"))

	err := client.Admin().Users().Unsuspend("83sv4lyx22")
	if err != nil {
		log.Printf("[Admin/Users] %s", err)

		return
	}
}
