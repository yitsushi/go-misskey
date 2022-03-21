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

func TestService_Create(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/admin/accounts/create",
		RequestData:  &users.CreateRequest{},
		ResponseFile: "create.json",
		StatusCode:   http.StatusOK,
	})

	user, err := client.Admin().Users().Create("sdk_test", "pass")
	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, "xxxxxxxx", user.Token)
	assert.Equal(t, "8f5irus8sp", user.ID)
}

func TestCreateRequest_Validate(t *testing.T) {
	test.ValidateRequests(
		t,
		[]core.BaseRequest{
			users.CreateRequest{},
			users.CreateRequest{Username: ""},
			users.CreateRequest{Password: ""},
			users.CreateRequest{Username: "", Password: ""},
			users.CreateRequest{Username: "asdsad", Password: ""},
			users.CreateRequest{Username: "", Password: "asdasd"},
		},
		[]core.BaseRequest{
			users.CreateRequest{Username: "sdk_test", Password: "pass"},
		},
	)
}

func ExampleService_Create() {
	client, _ := misskey.NewClientWithOptions(misskey.WithSimpleConfig("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN")))

	user, err := client.Admin().Users().Create("sdk_test", "pass")
	if err != nil {
		log.Printf("[Admin/Users] %s", err)

		return
	}

	log.Printf("[Admin/Users] %s", user.Username)
}
