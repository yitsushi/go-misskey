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

func TestService_Mentions(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/notes/mentions",
		RequestData:  &notes.MentionsRequest{},
		ResponseFile: "mentions.json",
		StatusCode:   http.StatusOK,
	})

	notes, err := client.Notes().Mentions(notes.MentionsRequest{
		Limit: 3,
	})
	if !assert.NoError(t, err) {
		return
	}

	assert.Len(t, notes, 3)
}

func TestMentionsRequest_Validate(t *testing.T) {
	test.ValidateRequests(
		t,
		[]core.BaseRequest{
			notes.MentionsRequest{},
		},
		[]core.BaseRequest{
			notes.MentionsRequest{
				Limit: 3,
			},
		},
	)
}

func ExampleService_Mentions() {
	client, _ := misskey.NewClientWithOptions(misskey.WithSimpleConfig("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN")))
	client.LogLevel(logrus.DebugLevel)

	noteList, err := client.Notes().Mentions(notes.MentionsRequest{
		Limit: 3,
	})
	if err != nil {
		log.Printf("[Notes] Error happened: %s", err)

		return
	}

	for _, note := range noteList {
		log.Println(note.Text)
	}
}
