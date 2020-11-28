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

func TestService_Show(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/admin/show-user",
		RequestData:  &users.ShowRequest{},
		ResponseFile: "show.json",
		StatusCode:   http.StatusOK,
	})

	user, err := client.Admin().Users().Show("83sv4lyx22")
	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, "efertone", user.Username)
	assert.Len(t, user.Emojis, 3)
}

func TestShowRequest_Validate(t *testing.T) {
	test.ValidateRequests(
		t,
		[]core.BaseRequest{
			users.ShowRequest{},
		},
		[]core.BaseRequest{
			users.ShowRequest{UserID: "asd"},
		},
	)
}

func ExampleService_Show() {
	client := misskey.NewClient("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN"))

	user, err := client.Admin().Users().Show("83sv4lyx22")
	if err != nil {
		log.Printf("[Admin/Users] %s", err)

		return
	}

	log.Printf("[Admin/Users] %s", user.Username)
}
