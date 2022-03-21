package emoji_test

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/services/admin/emoji"
	"github.com/yitsushi/go-misskey/test"
)

func TestService_Remove(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/admin/emoji/remove",
		RequestData:  &emoji.RemoveRequest{},
		ResponseFile: "empty",
		StatusCode:   http.StatusNoContent,
	})

	err := client.Admin().Emoji().Remove("8fbtx0k2ok")

	assert.NoError(t, err)
}

func TestRemoveRequest_Validate(t *testing.T) {
	test.ValidateRequests(
		t,
		[]core.BaseRequest{
			emoji.RemoveRequest{},
		},
		[]core.BaseRequest{
			emoji.RemoveRequest{ID: "asd"},
		},
	)
}

func ExampleService_Remove() {
	client, _ := misskey.NewClientWithOptions(misskey.WithSimpleConfig("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN")))

	err := client.Admin().Emoji().Remove("8fbtx0k2ok")
	if err != nil {
		log.Printf("[Admin/Emoji] %s", err)
	}
}
