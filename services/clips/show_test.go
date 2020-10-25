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

func TestService_Show(t *testing.T) {
	mockClient := test.SimpleMockEndpoint(&test.SimpleMockOptions{
		Endpoint:     "/api/clips/show",
		RequestData:  &clips.ShowRequest{},
		ResponseFile: "show.json",
		StatusCode:   http.StatusOK,
	})

	client := misskey.NewClient("https://localhost", "thisistoken")
	client.HTTPClient = mockClient

	clip, err := client.Clips().Show(clips.ShowRequest{
		ClipID: "8drxu3ckca",
	})
	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, "test clip", clip.Name)
}

func TestShowRequest_Validate(t *testing.T) {
	testCase := clips.ShowRequest{}
	assert.Error(t, testCase.Validate())
}

func ExampleService_Show() {
	client := misskey.NewClient("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN"))

	clip, err := client.Clips().Show(clips.ShowRequest{
		ClipID: "8drxu3ckca",
	})
	if err != nil {
		log.Printf("[Clips/Show] %s", err)

		return
	}

	log.Printf("[Clips/Show] %s", clip.Name)
}
