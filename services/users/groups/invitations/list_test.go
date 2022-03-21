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

func TestService_List(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/i/user-group-invites",
		RequestData:  &invitations.ListRequest{},
		ResponseFile: "list.json",
		StatusCode:   http.StatusOK,
	})

	invs, err := client.Users().Groups().Invitations().List(invitations.ListRequest{})
	if !assert.NoError(t, err) {
		return
	}

	assert.Len(t, invs, 1)
}

func TestListRequest_Validate(t *testing.T) {
	test.ValidateRequests(
		t,
		[]core.BaseRequest{},
		[]core.BaseRequest{
			invitations.ListRequest{},
		},
	)
}

func ExampleService_List() {
	client, _ := misskey.NewClientWithOptions(misskey.WithSimpleConfig("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN")))

	invs, err := client.Users().Groups().Invitations().List(invitations.ListRequest{})
	if err != nil {
		log.Printf("[Users/Groups/Invitations/List] %s", err)

		return
	}

	for _, inv := range invs {
		log.Printf("[Users/Groups/Invitations/List] <%s> %s", inv.ID, inv.Group.Name)
	}
}
