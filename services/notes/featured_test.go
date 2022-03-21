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

func TestService_Featured(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/notes/featured",
		RequestData:  &notes.FeaturedRequest{},
		ResponseFile: "featured.json",
		StatusCode:   http.StatusOK,
	})

	noteList, err := client.Notes().Featured(notes.FeaturedRequest{
		Limit: 10,
	})
	if !assert.NoError(t, err) {
		return
	}

	assert.Len(t, noteList, 1)
	assert.Equal(t, "efertone", noteList[0].User.Username)
}

func TestFeaturedRequest_Validate(t *testing.T) {
	test.ValidateRequests(
		t,
		[]core.BaseRequest{
			notes.FeaturedRequest{},
		},
		[]core.BaseRequest{
			notes.FeaturedRequest{Limit: 20},
		},
	)
}

func ExampleService_Featured() {
	client, _ := misskey.NewClientWithOptions(misskey.WithSimpleConfig("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN")))
	client.LogLevel(logrus.DebugLevel)

	noteList, err := client.Notes().Featured(notes.FeaturedRequest{
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
