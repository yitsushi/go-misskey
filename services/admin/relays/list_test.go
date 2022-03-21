package relays_test

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/services/admin/relays"
	"github.com/yitsushi/go-misskey/test"
)

func TestService_List(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/admin/relays/list",
		RequestData:  &relays.ListRequest{},
		ResponseFile: "list.json",
		StatusCode:   http.StatusOK,
	})

	relays, err := client.Admin().Relays().List()

	assert.NoError(t, err)
	assert.Len(t, relays, 3)
}

func TestListRequest_Validate(t *testing.T) {
	test.ValidateRequests(
		t,
		[]core.BaseRequest{},
		[]core.BaseRequest{
			relays.ListRequest{},
		},
	)
}

func ExampleService_List() {
	client := misskey.NewClient("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN"))

	relays, err := client.Admin().Relays().List()
	if err != nil {
		log.Printf("[Admin/Relays/List] %s", err)

		return
	}

	for _, relay := range relays {
		log.Printf("[Admin/Relays/List] <%s> %s", relay.Status, relay.Inbox)
	}
}
