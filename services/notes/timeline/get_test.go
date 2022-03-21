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

func TestService_Get(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/notes/timeline",
		RequestData:  &timeline.GetRequest{},
		ResponseFile: "get.json",
		StatusCode:   http.StatusOK,
	})

	noteList, err := client.Notes().Timeline().Get(timeline.GetRequest{
		Limit: 3,
	})
	if !assert.NoError(t, err) {
		return
	}

	assert.Len(t, noteList, 3)
	assert.Equal(t, "aoife", noteList[0].User.Username)
}

func TestGetRequest_Validate(t *testing.T) {
	test.ValidateRequests(
		t,
		[]core.BaseRequest{
			timeline.GetRequest{},
			timeline.GetRequest{Limit: 3000},
		},
		[]core.BaseRequest{
			timeline.GetRequest{Limit: 20},
		},
	)
}

func ExampleService_Get() {
	client, _ := misskey.NewClientWithOptions(misskey.WithSimpleConfig("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN")))
	client.LogLevel(logrus.DebugLevel)

	noteList, err := client.Notes().Timeline().Get(timeline.GetRequest{
		Limit: 10,
	})
	if err != nil {
		log.Printf("[Notes] Error happened: %s", err)

		return
	}

	for _, note := range noteList {
		log.Printf(" - %s", note.Text)
	}
}
