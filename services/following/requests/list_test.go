package requests_test

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/services/following/requests"
	"github.com/yitsushi/go-misskey/test"
)

func TestService_List(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/following/requests/list",
		RequestData:  &requests.ListRequest{},
		ResponseFile: "list.json",
		StatusCode:   http.StatusOK,
	})

	list, err := client.Following().Requests().List()
	if !assert.NoError(t, err) {
		return
	}

	assert.Len(t, list, 1)
	assert.Equal(t, "efertone", list[0].Follower.Username)
	assert.Equal(t, "kiki_test", list[0].Followee.Username)
}

func TestListRequest_Validate(t *testing.T) {
	test.ValidateRequests(
		t,
		[]core.BaseRequest{},
		[]core.BaseRequest{
			requests.ListRequest{},
		},
	)
}

func ExampleService_List() {
	client, _ := misskey.NewClientWithOptions(misskey.WithSimpleConfig("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN")))

	list, err := client.Following().Requests().List()
	if err != nil {
		log.Printf("[Following/Requests/List] %s", err)

		return
	}

	for _, item := range list {
		log.Printf(
			"[Following/Requests/List] %s -> %s",
			item.Follower.Username,
			item.Followee.Username,
		)
	}
}
