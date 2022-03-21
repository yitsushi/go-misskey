package timeline_test

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/services/notes/timeline"
	"github.com/yitsushi/go-misskey/test"
)

func TestService_UserList(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/notes/user-list-timeline",
		RequestData:  &timeline.UserListRequest{},
		ResponseFile: "get.json",
		StatusCode:   http.StatusOK,
	})

	noteList, err := client.Notes().Timeline().UserList(timeline.UserListRequest{
		Limit:  3,
		ListID: "listid",
	})
	if !assert.NoError(t, err) {
		return
	}

	assert.Len(t, noteList, 3)
	assert.Equal(t, "aoife", noteList[0].User.Username)
}

func TestUserListRequest_Validate(t *testing.T) {
	test.ValidateRequests(
		t,
		[]core.BaseRequest{
			timeline.UserListRequest{},
			timeline.UserListRequest{ListID: "listid"},
			timeline.UserListRequest{Limit: 3000},
			timeline.UserListRequest{ListID: "listid", Limit: 3000},
		},
		[]core.BaseRequest{
			timeline.UserListRequest{ListID: "listid", Limit: 20},
		},
	)
}

func ExampleService_UserList() {
	client, _ := misskey.NewClientWithOptions(misskey.WithSimpleConfig("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN")))
	client.LogLevel(logrus.DebugLevel)

	noteList, err := client.Notes().Timeline().UserList(timeline.UserListRequest{
		Limit:  10,
		ListID: "listid",
	})
	if err != nil {
		log.Printf("[Notes] Error happened: %s", err)

		return
	}

	for _, note := range noteList {
		log.Printf(" - %s", note.Text)
	}
}
