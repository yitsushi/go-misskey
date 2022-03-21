package clips_test

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/services/clips"
	"github.com/yitsushi/go-misskey/test"
)

func TestService_List(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/clips/list",
		RequestData:  &clips.ListRequest{},
		ResponseFile: "list.json",
		StatusCode:   http.StatusOK,
	})

	clips, err := client.Clips().List()
	if !assert.NoError(t, err) {
		return
	}

	assert.Len(t, clips, 1)
}

func ExampleService_List() {
	client, _ := misskey.NewClientWithOptions(misskey.WithSimpleConfig("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN")))

	clips, err := client.Clips().List()
	if err != nil {
		log.Printf("[Clips/List] %s", err)

		return
	}

	for _, clip := range clips {
		log.Printf("[Clips/List] %s", clip.Name)
	}
}
