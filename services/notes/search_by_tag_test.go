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

func TestService_SearchByTag(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/notes/search-by-tag",
		RequestData:  &notes.SearchByTagRequest{},
		ResponseFile: "search.json",
		StatusCode:   http.StatusOK,
	})

	noteList, err := client.Notes().SearchByTag(notes.SearchByTagRequest{
		Tag:   "test",
		Limit: 2,
	})
	if !assert.NoError(t, err) {
		return
	}

	assert.Len(t, noteList, 2)
}

func TestSearchByTagRequest_Validate(t *testing.T) {
	test.ValidateRequests(
		t,
		[]core.BaseRequest{
			notes.SearchByTagRequest{},
			notes.SearchByTagRequest{Tag: ""},
			notes.SearchByTagRequest{Query: [][]string{}},
			notes.SearchByTagRequest{Limit: 10},
		},
		[]core.BaseRequest{
			notes.SearchByTagRequest{Tag: "test", Limit: 20},
		},
	)
}

func ExampleService_SearchByTag() {
	client, _ := misskey.NewClientWithOptions(misskey.WithSimpleConfig("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN")))
	client.LogLevel(logrus.DebugLevel)

	noteList, err := client.Notes().SearchByTag(notes.SearchByTagRequest{
		Tag:   "test",
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
