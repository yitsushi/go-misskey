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

func TestService_Joined(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/users/groups/joined",
		RequestData:  &groups.JoinedRequest{},
		ResponseFile: "joined.json",
		StatusCode:   http.StatusOK,
	})

	resp, err := client.Users().Groups().Joined()
	if !assert.NoError(t, err) {
		return
	}

	assert.Len(t, resp, 1)
}

func ExampleService_Joined() {
	client := misskey.NewClient("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN"))

	resp, err := client.Users().Groups().Joined()
	if err != nil {
		log.Printf("[Users/Groups/Joined] %s", err)

		return
	}

	for _, group := range resp {
		log.Printf("[Users/Groups/Joined] %s", group.Name)
	}
}
