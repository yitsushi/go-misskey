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

func TestService_Create(t *testing.T) {
	mockClient := test.SimpleMockEndpoint(&test.SimpleMockOptions{
		Endpoint:     "/api/clips/create",
		RequestData:  &clips.CreateRequest{},
		ResponseFile: "create.json",
		StatusCode:   http.StatusOK,
	})

	client := misskey.NewClient("https://localhost", "thisistoken")
	client.HTTPClient = mockClient

	clip, err := client.Clips().Create(clips.CreateRequest{
		Name: "test clip",
	})
	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, "test clip", clip.Name)
}

func TestCreateRequest_Validate(t *testing.T) {
	testCase := clips.CreateRequest{}
	assert.Error(t, testCase.Validate())

	testCase = clips.CreateRequest{
		Name: strings.Repeat("a", 101),
	}
	assert.Error(t, testCase.Validate())
}

func ExampleService_Create() {
	client := misskey.NewClient("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN"))

	clip, err := client.Clips().Create(clips.CreateRequest{
		Name: "test clip",
	})
	if err != nil {
		log.Printf("[Clips/Create] %s", err)

		return
	}

	log.Printf("[Clips/Create] %s created", clip.Name)
}
