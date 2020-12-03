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

func TestService_List(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/admin/emoji/list",
		RequestData:  &emoji.ListRequest{},
		ResponseFile: "list.json",
		StatusCode:   http.StatusOK,
	})

	response, err := client.Admin().Emoji().List(emoji.ListRequest{})

	if !assert.NoError(t, err) {
		return
	}

	assert.Len(t, response, 10)
}

func TestListRequest_Validate(t *testing.T) {
	test.ValidateRequests(
		t,
		[]core.BaseRequest{},
		[]core.BaseRequest{
			emoji.ListRequest{},
		},
	)
}

func ExampleService_List() {
	client := misskey.NewClient("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN"))

	response, err := client.Admin().Emoji().List(emoji.ListRequest{})
	if err != nil {
		log.Printf("[Admin/Emoji] %s", err)

		return
	}

	for _, item := range response {
		log.Printf("[Admin/Emoji] %s", core.StringValue(item.Name))
	}
}
