package clips_test

import (
	"log"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/services/clips"
	"github.com/yitsushi/go-misskey/test"
)

func TestService_Update(t *testing.T) {
	mockClient := test.SimpleMockEndpoint(&test.SimpleMockOptions{
		Endpoint:     "/api/clips/update",
		RequestData:  &clips.UpdateRequest{},
		ResponseFile: "update.json",
		StatusCode:   http.StatusOK,
	})

	client := misskey.NewClient("https://localhost", "thisistoken")
	client.HTTPClient = mockClient

	clip, err := client.Clips().Update(clips.UpdateRequest{
		ClipID: "8drxu3ckca",
		Name:   "new test",
	})
	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, "new test", clip.Name)
}

func TestUpdateRequest_Validate(t *testing.T) {
	testCase := clips.UpdateRequest{}
	assert.Error(t, testCase.Validate())

	testCase = clips.UpdateRequest{
		ClipID: "8drxu3ckca",
	}
	assert.Error(t, testCase.Validate())

	testCase = clips.UpdateRequest{
		ClipID: "8drxu3ckca",
		Name:   strings.Repeat("a", 101),
	}
	assert.Error(t, testCase.Validate())
}

func ExampleService_Update() {
	client := misskey.NewClient("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN"))

	clip, err := client.Clips().Update(clips.UpdateRequest{
		ClipID: "8drxu3ckca",
		Name:   "new test",
	})
	if err != nil {
		log.Printf("[Clips/Update] %s", err)

		return
	}

	log.Printf("[Clips/Update] %s updated", clip.Name)
}
