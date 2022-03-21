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

func TestService_Search(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/notes/search",
		RequestData:  &notes.SearchRequest{},
		ResponseFile: "search.json",
		StatusCode:   http.StatusOK,
	})

	noteList, err := client.Notes().Search(notes.SearchRequest{
		Query: "Winter first movement",
		Limit: 10,
	})
	if !assert.NoError(t, err) {
		return
	}

	assert.Len(t, noteList, 2)
	assert.Equal(t, "efertone", noteList[0].User.Username)
}

func TestSearchRequest_Validate(t *testing.T) {
	test.ValidateRequests(
		t,
		[]core.BaseRequest{
			notes.SearchRequest{},
			notes.SearchRequest{Query: "asd"},
		},
		[]core.BaseRequest{
			notes.SearchRequest{Query: "asd", Limit: 20},
		},
	)
}

func ExampleService_Search() {
	client, _ := misskey.NewClientWithOptions(misskey.WithSimpleConfig("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN")))
	client.LogLevel(logrus.DebugLevel)

	noteList, err := client.Notes().Search(notes.SearchRequest{
		Query: "hack%",
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
