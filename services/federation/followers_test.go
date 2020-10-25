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

func TestService_Followers(t *testing.T) {
	mockClient := test.SimpleMockEndpoint(&test.SimpleMockOptions{
		Endpoint:     "/api/federation/followers",
		RequestData:  &federation.FollowersRequest{},
		ResponseFile: "followers.json",
		StatusCode:   http.StatusOK,
	})

	client := misskey.NewClient("https://localhost", "thisistoken")
	client.HTTPClient = mockClient

	followers, err := client.Federation().Followers(federation.FollowersRequest{
		Limit: 10,
		Host:  "slippy.xyz",
	})
	if !assert.NoError(t, err) {
		return
	}

	assert.Len(t, followers, 2)
}

func TestFollowersRequest_Validate(t *testing.T) {
	testCase := federation.FollowersRequest{}
	assert.Error(t, testCase.Validate())

	testCase = federation.FollowersRequest{
		Host: "slippy.xyz",
	}
	assert.Error(t, testCase.Validate())

	testCase = federation.FollowersRequest{
		Host:  "slippy.xyz",
		Limit: 10,
	}
	assert.NoError(t, testCase.Validate())
}

func ExampleService_Followers() {
	client := misskey.NewClient("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN"))

	resp, err := client.Federation().Followers(federation.FollowersRequest{
		Limit: 10,
		Host:  "slippy.xyz",
	})
	if err != nil {
		log.Printf("[Federation/Followers] %s", err)

		return
	}

	log.Printf("[Federation/Followers] %v listed", resp)
}
