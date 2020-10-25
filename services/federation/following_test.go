package federation_test

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/services/federation"
	"github.com/yitsushi/go-misskey/test"
)

func TestService_Following(t *testing.T) {
	mockClient := test.SimpleMockEndpoint(&test.SimpleMockOptions{
		Endpoint:     "/api/federation/following",
		RequestData:  &federation.FollowersRequest{},
		ResponseFile: "followers.json",
		StatusCode:   http.StatusOK,
	})

	client := misskey.NewClient("https://localhost", "thisistoken")
	client.HTTPClient = mockClient

	followers, err := client.Federation().Following(federation.FollowingRequest{
		Limit: 10,
		Host:  "slippy.xyz",
	})
	if !assert.NoError(t, err) {
		return
	}

	assert.Len(t, followers, 2)
}

func TestFollowingRequest_Validate(t *testing.T) {
	testCase := federation.FollowingRequest{}
	assert.Error(t, testCase.Validate())

	testCase = federation.FollowingRequest{
		Host: "slippy.xyz",
	}
	assert.Error(t, testCase.Validate())

	testCase = federation.FollowingRequest{
		Host:  "slippy.xyz",
		Limit: 10,
	}
	assert.NoError(t, testCase.Validate())
}

func ExampleService_Following() {
	client := misskey.NewClient("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN"))

	resp, err := client.Federation().Following(federation.FollowingRequest{
		Limit: 100,
		Host:  "slippy.xyz",
	})
	if err != nil {
		log.Printf("[Federation/Following] %s", err)

		return
	}

	log.Printf("[Federation/Following] %v listed", resp)
}
