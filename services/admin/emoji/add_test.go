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

func TestService_Add(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/admin/emoji/add",
		RequestData:  &emoji.AddRequest{},
		ResponseFile: "add.json",
		StatusCode:   http.StatusOK,
	})

	response, err := client.Admin().Emoji().Add(emoji.AddRequest{
		Name:   "Test",
		FileID: "8fbtx0k2ok",
	})

	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, "8fbu3oru8r", response)
}

func TestAddRequest_Validate(t *testing.T) {
	test.ValidateRequests(
		t,
		[]core.BaseRequest{
			emoji.AddRequest{},
			emoji.AddRequest{Name: ""},
			emoji.AddRequest{FileID: ""},
			emoji.AddRequest{Name: "", FileID: ""},
		},
		[]core.BaseRequest{
			emoji.AddRequest{
				Name:   "test",
				FileID: "8fbtx0k2ok",
			},
		},
	)
}

func ExampleService_Add() {
	client, _ := misskey.NewClientWithOptions(misskey.WithSimpleConfig("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN")))

	emojiID, err := client.Admin().Emoji().Add(emoji.AddRequest{
		Name:   "EmojiName",
		FileID: "8fbtx0k2ok",
	})
	if err != nil {
		log.Printf("[Admin/Emoji] %s", err)

		return
	}

	log.Printf("[Admin/Emoji] %s", emojiID)
}
