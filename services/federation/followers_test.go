package federation_test

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/services/federation"
	"github.com/yitsushi/go-misskey/test"
)

func TestService_Followers(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/federation/followers",
		RequestData:  &federation.FollowersRequest{},
		ResponseFile: "followers.json",
		StatusCode:   http.StatusOK,
	})

	followers, err := client.Federation().Followers(federation.FollowersRequest{
		Limit: 2,
		Host:  "slippy.xyz",
	})
	if !assert.NoError(t, err) {
		return
	}

	assert.Len(t, followers, 2)

	assert.Equal(t, "84l9tncinl", followers[0].FolloweeID)
	assert.Equal(t, "83sv4lyx22", followers[0].FollowerID)
}

func TestFollowersRequest_Validate(t *testing.T) {
	test.ValidateRequests(
		t,
		[]core.BaseRequest{
			federation.FollowersRequest{},
			federation.FollowersRequest{Host: "slippy.xyz"},
		},
		[]core.BaseRequest{},
	)
}

func ExampleService_Followers() {
	client, _ := misskey.NewClientWithOptions(misskey.WithSimpleConfig("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN")))

	resp, err := client.Federation().Followers(federation.FollowersRequest{
		Limit: 40,
		Host:  "slippy.xyz",
	})
	if err != nil {
		log.Printf("[Federation/Followers] %s", err)

		return
	}

	log.Printf("[Federation/Followers] %v listed", resp)
}
