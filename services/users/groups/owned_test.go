package groups_test

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/services/users/groups"
	"github.com/yitsushi/go-misskey/test"
)

func TestService_Owned(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/users/groups/owned",
		RequestData:  &groups.OwnedRequest{},
		ResponseFile: "owned.json",
		StatusCode:   http.StatusOK,
	})

	resp, err := client.Users().Groups().Owned()
	if !assert.NoError(t, err) {
		return
	}

	assert.Len(t, resp, 1)
}

func ExampleService_Owned() {
	client, _ := misskey.NewClientWithOptions(misskey.WithSimpleConfig("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN")))

	resp, err := client.Users().Groups().Owned()
	if err != nil {
		log.Printf("[Users/Groups/Owned] %s", err)

		return
	}

	for _, group := range resp {
		log.Printf("[Users/Groups/Owned] %s", group.Name)
	}
}
