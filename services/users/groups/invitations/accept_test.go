package invitations_test

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/services/users/groups/invitations"
	"github.com/yitsushi/go-misskey/test"
)

func TestService_Accept(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/users/groups/invitations/accept",
		RequestData:  &invitations.AcceptRequest{},
		ResponseFile: "empty",
		StatusCode:   http.StatusNoContent,
	})

	err := client.Users().Groups().Invitations().Accept("8y4nwgla5f")
	if !assert.NoError(t, err) {
		return
	}
}

func TestAcceptRequest_Validate(t *testing.T) {
	test.ValidateRequests(
		t,
		[]core.BaseRequest{
			invitations.AcceptRequest{},
		},
		[]core.BaseRequest{
			invitations.AcceptRequest{InvitationID: "8y4nwgla5f"},
		},
	)
}

func ExampleService_Accept() {
	client, _ := misskey.NewClientWithOptions(misskey.WithSimpleConfig("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN")))

	err := client.Users().Groups().Invitations().Accept("8y4nwgla5f")
	if err != nil {
		log.Printf("[Users/Groups/Invitations/Accept] %s", err)

		return
	}
}
