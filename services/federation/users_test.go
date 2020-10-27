package federation_test

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/services/federation"
	"github.com/yitsushi/go-misskey/test"
)

func TestService_Users(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/federation/users",
		RequestData:  &federation.UsersRequest{},
		ResponseFile: "users.json",
		StatusCode:   http.StatusOK,
	})

	users, err := client.Federation().Users(federation.UsersRequest{
		Limit: 2,
		Host:  "misskey.io",
	})
	if !assert.NoError(t, err) {
		return
	}

	assert.Len(t, users, 10)
}

func TestUsersRequest_Validate(t *testing.T) {
	test.ValidateRequests(
		t,
		[]core.BaseRequest{
			federation.UsersRequest{},
			federation.UsersRequest{Host: "misskey.io"},
		},
		[]core.BaseRequest{},
	)
}

func ExampleService_Users() {
	client := misskey.NewClient("https://misskey.io", os.Getenv("MISSKEY_TOKEN"))

	resp, err := client.Federation().Users(federation.UsersRequest{
		Limit: 90,
		Host:  "misskey.io",
	})
	if err != nil {
		log.Printf("[Federation/Users] %s", err)

		return
	}

	log.Printf("[Federation/Users] %v listed", resp)
}
