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

func TestService_Following(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/federation/following",
		RequestData:  &federation.FollowersRequest{},
		ResponseFile: "following.json",
		StatusCode:   http.StatusOK,
	})

	following, err := client.Federation().Following(federation.FollowingRequest{
		Limit: 1,
		Host:  "slippy.xyz",
	})
	if !assert.NoError(t, err) {
		return
	}

	assert.Len(t, following, 1)
	assert.Equal(t, "84wi1id9ev", following[0].FollowerID)
}

func TestFollowingRequest_Validate(t *testing.T) {
	test.ValidateRequests(
		t,
		[]core.BaseRequest{
			federation.FollowingRequest{},
			federation.FollowingRequest{Host: "slippy.xyz"},
		},
		[]core.BaseRequest{},
	)
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
