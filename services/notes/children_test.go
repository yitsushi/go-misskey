package notes_test

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/services/notes"
	"github.com/yitsushi/go-misskey/test"
)

func TestService_Children(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/notes/children",
		RequestData:  &notes.ChildrenRequest{},
		ResponseFile: "children.json",
		StatusCode:   http.StatusOK,
	})

	noteList, err := client.Notes().Children(notes.ChildrenRequest{
		NoteID: "noteid",
		Limit:  1,
	})
	if !assert.NoError(t, err) {
		return
	}

	assert.Len(t, noteList, 1)
}

func TestChildrenRequest_Validate(t *testing.T) {
	test.ValidateRequests(
		t,
		[]core.BaseRequest{
			notes.ChildrenRequest{},
			notes.ChildrenRequest{NoteID: ""},
			notes.ChildrenRequest{NoteID: "noteid"},
			notes.ChildrenRequest{Limit: 10},
			notes.ChildrenRequest{NoteID: "noteid", Limit: 200},
		},
		[]core.BaseRequest{
			notes.ChildrenRequest{NoteID: "noteid", Limit: 10},
		},
	)
}

func ExampleService_Children() {
	client, _ := misskey.NewClientWithOptions(misskey.WithSimpleConfig("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN")))
	client.LogLevel(logrus.DebugLevel)

	noteList, err := client.Notes().Children(notes.ChildrenRequest{
		NoteID: "noteid",
		Limit:  10,
	})
	if err != nil {
		log.Printf("[Notes] Error happened: %s", err)

		return
	}

	for _, note := range noteList {
		log.Printf(" - %s", note.Text)
	}
}
