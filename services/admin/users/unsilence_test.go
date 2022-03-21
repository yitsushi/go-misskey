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

func TestService_Unsilence(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/admin/unsilence-user",
		RequestData:  &users.UnsilenceRequest{},
		ResponseFile: "empty",
		StatusCode:   http.StatusNoContent,
	})

	err := client.Admin().Users().Unsilence("83sv4lyx22")
	assert.NoError(t, err)
}

func TestUnsilenceRequest_Validate(t *testing.T) {
	test.ValidateRequests(
		t,
		[]core.BaseRequest{
			users.UnsilenceRequest{},
		},
		[]core.BaseRequest{
			users.UnsilenceRequest{UserID: "asd"},
		},
	)
}

func ExampleService_Unsilence() {
	client, _ := misskey.NewClientWithOptions(misskey.WithSimpleConfig("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN")))

	err := client.Admin().Users().Unsilence("83sv4lyx22")
	if err != nil {
		log.Printf("[Admin/Users] %s", err)

		return
	}
}
