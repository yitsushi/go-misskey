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

func TestService_Show(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/users/groups/show",
		RequestData:  &groups.ShowRequest{},
		ResponseFile: "show.json",
		StatusCode:   http.StatusOK,
	})

	group, err := client.Users().Groups().Show("93tyd132e7")
	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, "Test", group.Name)
}

func TestShowRequest_Validate(t *testing.T) {
	test.ValidateRequests(
		t,
		[]core.BaseRequest{
			groups.ShowRequest{},
		},
		[]core.BaseRequest{
			groups.ShowRequest{GroupID: "93tyd132e7"},
		},
	)
}

func ExampleService_Show() {
	client, _ := misskey.NewClientWithOptions(misskey.WithSimpleConfig("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN")))

	group, err := client.Users().Groups().Show("93tyd132e7")
	if err != nil {
		log.Printf("[Users/Groups/Show] %s", err)

		return
	}

	log.Printf("[Users/Groups/Show] %s", group.Name)
}
