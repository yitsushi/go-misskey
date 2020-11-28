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

func TestService_List(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/admin/show-users",
		RequestData:  &users.ListRequest{},
		ResponseFile: "list.json",
		StatusCode:   http.StatusOK,
	})

	userList, err := client.Admin().Users().List(users.ListRequest{
		Limit: 4,
	})
	if !assert.NoError(t, err) {
		return
	}

	assert.Len(t, userList, 4)
}

func TestListRequest_Validate(t *testing.T) {
	test.ValidateRequests(
		t,
		[]core.BaseRequest{},
		[]core.BaseRequest{
			users.ListRequest{},
		},
	)
}

func ExampleService_List() {
	client := misskey.NewClient("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN"))

	userList, err := client.Admin().Users().List(users.ListRequest{
		Limit: 4,
	})
	if err != nil {
		log.Printf("[Admin/Users] %s", err)

		return
	}

	for _, user := range userList {
		log.Printf("[Admin/Users] %s", user.Username)
	}
}
