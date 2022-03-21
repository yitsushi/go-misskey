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

func TestService_Copy(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/admin/emoji/copy",
		RequestData:  &emoji.CopyRequest{},
		ResponseFile: "copy.json",
		StatusCode:   http.StatusOK,
	})

	response, err := client.Admin().Emoji().Copy(emoji.CopyRequest{
		EmojiID: "8fbtx0k2ok",
	})

	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, "8fbu3oru8r", response)
}

func TestCopyRequest_Validate(t *testing.T) {
	test.ValidateRequests(
		t,
		[]core.BaseRequest{
			emoji.CopyRequest{},
		},
		[]core.BaseRequest{
			emoji.CopyRequest{EmojiID: "asd"},
		},
	)
}

func ExampleService_Copy() {
	client, _ := misskey.NewClientWithOptions(misskey.WithSimpleConfig("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN")))

	emojiID, err := client.Admin().Emoji().Copy(emoji.CopyRequest{
		EmojiID: "8fbtx0k2ok",
	})
	if err != nil {
		log.Printf("[Admin/Emoji] %s", err)

		return
	}

	log.Printf("[Admin/Emoji] %s", emojiID)
}
