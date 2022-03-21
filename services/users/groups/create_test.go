package groups_test

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/services/users/groups"
	"github.com/yitsushi/go-misskey/test"
)

func TestService_Create(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/users/groups/create",
		RequestData:  &groups.CreateRequest{},
		ResponseFile: "show.json",
		StatusCode:   http.StatusOK,
	})

	group, err := client.Users().Groups().Create("Test")
	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, "Test", group.Name)
}

func TestCreateRequest_Validate(t *testing.T) {
	test.ValidateRequests(
		t,
		[]core.BaseRequest{
			groups.CreateRequest{},
		},
		[]core.BaseRequest{
			groups.CreateRequest{Name: "Test"},
		},
	)
}

func ExampleService_Create() {
	client := misskey.NewClient("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN"))

	group, err := client.Users().Groups().Create("Test")
	if err != nil {
		log.Printf("[Users/Groups/Create] %s", err)

		return
	}

	log.Printf("[Users/Groups/Create] %s", group.Name)
}
