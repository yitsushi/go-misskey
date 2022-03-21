package groups_test

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/services/users/groups"
	"github.com/yitsushi/go-misskey/test"
)

func TestService_Kick(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/users/groups/pull",
		RequestData:  &groups.KickRequest{},
		ResponseFile: "empty",
		StatusCode:   http.StatusNoContent,
	})

	err := client.Users().Groups().Kick("93tyd132e7", "83sv4lyx22")
	if !assert.NoError(t, err) {
		return
	}
}

func TestKickRequest_Validate(t *testing.T) {
	test.ValidateRequests(
		t,
		[]core.BaseRequest{
			groups.KickRequest{},
			groups.KickRequest{GroupID: "93tyd132e7"},
			groups.KickRequest{UserID: "93tyd132e7"},
		},
		[]core.BaseRequest{
			groups.KickRequest{GroupID: "93tyd132e7", UserID: "83sv4lyx22"},
		},
	)
}

func ExampleService_Kick() {
	client, _ := misskey.NewClientWithOptions(misskey.WithSimpleConfig("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN")))

	err := client.Users().Groups().Kick("93tyd132e7", "83sv4lyx22")
	if err != nil {
		log.Printf("[Users/Groups/Kick] %s", err)

		return
	}
}
