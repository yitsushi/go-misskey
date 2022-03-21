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

func TestService_Remove(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/admin/relays/remove",
		RequestData:  &relays.RemoveRequest{},
		ResponseFile: "empty",
		StatusCode:   http.StatusNoContent,
	})

	err := client.Admin().Relays().Remove("https://something.tld/inbox")

	assert.NoError(t, err)
}

func TestRemoveRequest_Validate(t *testing.T) {
	test.ValidateRequests(
		t,
		[]core.BaseRequest{
			relays.RemoveRequest{},
		},
		[]core.BaseRequest{
			relays.RemoveRequest{Inbox: "https://something.tld/inbox"},
		},
	)
}

func ExampleService_Remove() {
	client := misskey.NewClient("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN"))

	err := client.Admin().Relays().Remove("https://something.tld/inbox")
	if err != nil {
		log.Printf("[Admin/Relays/Remove] %s", err)

		return
	}
}
