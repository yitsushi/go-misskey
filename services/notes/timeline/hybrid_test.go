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

func TestService_Hybrid(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/notes/hybrid-timeline",
		RequestData:  &timeline.HybridRequest{},
		ResponseFile: "get.json",
		StatusCode:   http.StatusOK,
	})

	noteList, err := client.Notes().Timeline().Hybrid(timeline.HybridRequest{
		Limit: 3,
	})
	if !assert.NoError(t, err) {
		return
	}

	assert.Len(t, noteList, 3)
	assert.Equal(t, "aoife", noteList[0].User.Username)
}

func TestHybridRequest_Validate(t *testing.T) {
	test.ValidateRequests(
		t,
		[]core.BaseRequest{
			timeline.HybridRequest{},
			timeline.HybridRequest{Limit: 3000},
		},
		[]core.BaseRequest{
			timeline.HybridRequest{Limit: 20},
		},
	)
}

func ExampleService_Hybrid() {
	client, _ := misskey.NewClientWithOptions(misskey.WithSimpleConfig("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN")))
	client.LogLevel(logrus.DebugLevel)

	noteList, err := client.Notes().Timeline().Hybrid(timeline.HybridRequest{
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
