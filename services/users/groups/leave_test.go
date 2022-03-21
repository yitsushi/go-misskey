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

func TestService_Leave(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/users/groups/leave",
		RequestData:  &groups.LeaveRequest{},
		ResponseFile: "empty",
		StatusCode:   http.StatusNoContent,
	})

	err := client.Users().Groups().Leave("93tyd132e7")
	if !assert.NoError(t, err) {
		return
	}
}

func TestLeaveRequest_Validate(t *testing.T) {
	test.ValidateRequests(
		t,
		[]core.BaseRequest{
			groups.LeaveRequest{},
		},
		[]core.BaseRequest{
			groups.LeaveRequest{GroupID: "93tyd132e7"},
		},
	)
}

func ExampleService_Leave() {
	client, _ := misskey.NewClientWithOptions(misskey.WithSimpleConfig("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN")))

	err := client.Users().Groups().Leave("93tyd132e7")
	if err != nil {
		log.Printf("[Users/Groups/Leave] %s", err)

		return
	}
}
