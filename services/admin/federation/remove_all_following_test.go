package federation_test

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/services/admin/federation"
	"github.com/yitsushi/go-misskey/test"
)

func TestService_RemoveAllFollowing(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/admin/federation/remove-all-following",
		RequestData:  &federation.RemoveAllFollowingRequest{},
		ResponseFile: "empty",
		StatusCode:   http.StatusNoContent,
	})

	err := client.Admin().Federation().RemoveAllFollowing(federation.RemoveAllFollowingRequest{
		Host: "quey.org",
	})

	assert.NoError(t, err)
}

func TestRemoveAllFollowingRequest_Validate(t *testing.T) {
	test.ValidateRequests(
		t,
		[]core.BaseRequest{
			federation.RemoveAllFollowingRequest{},
		},
		[]core.BaseRequest{
			federation.RemoveAllFollowingRequest{
				Host: "quey.org",
			},
		},
	)
}

func ExampleService_RemoveAllFollowing() {
	client := misskey.NewClient("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN"))

	err := client.Admin().Federation().RemoveAllFollowing(federation.RemoveAllFollowingRequest{
		Host: "quey.org",
	})
	if err != nil {
		log.Printf("[Admin/Federation] %s", err)
	}
}
