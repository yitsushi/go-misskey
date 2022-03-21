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

func TestService_Update(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/admin/emoji/update",
		RequestData:  &emoji.UpdateRequest{},
		ResponseFile: "empty",
		StatusCode:   http.StatusNoContent,
	})

	err := client.Admin().Emoji().Update(emoji.UpdateRequest{
		ID:   "8fbtx0k2ok",
		Name: "name",
	})

	assert.NoError(t, err)
}

func TestUpdateRequest_Validate(t *testing.T) {
	test.ValidateRequests(
		t,
		[]core.BaseRequest{
			emoji.UpdateRequest{},
			emoji.UpdateRequest{ID: "asd"},
		},
		[]core.BaseRequest{
			emoji.UpdateRequest{ID: "asd", Name: "name"},
			emoji.UpdateRequest{
				ID:       "asd",
				Name:     "name",
				Category: "cat",
			},
			emoji.UpdateRequest{
				ID:       "asd",
				Name:     "name",
				Category: "cat",
				Aliases:  []string{"alias1"},
			},
		},
	)
}

func ExampleService_Update() {
	client, _ := misskey.NewClientWithOptions(misskey.WithSimpleConfig("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN")))

	err := client.Admin().Emoji().Update(emoji.UpdateRequest{
		ID:       "8fbtx0k2ok",
		Name:     "emoji-name",
		Category: "cat",
	})
	if err != nil {
		log.Printf("[Admin/Emoji] %s", err)
	}
}
