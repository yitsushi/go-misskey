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

func TestService_Conversation(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/notes/conversation",
		RequestData:  &notes.ConversationRequest{},
		ResponseFile: "conversation.json",
		StatusCode:   http.StatusOK,
	})

	noteList, err := client.Notes().Conversation(notes.ConversationRequest{
		NoteID: "noteid",
		Limit:  10,
	})
	if !assert.NoError(t, err) {
		return
	}

	assert.Len(t, noteList, 4)
}

func TestConversationRequest_Validate(t *testing.T) {
	test.ValidateRequests(
		t,
		[]core.BaseRequest{
			notes.ConversationRequest{},
			notes.ConversationRequest{NoteID: "asd"},
			notes.ConversationRequest{Limit: 10},
		},
		[]core.BaseRequest{
			notes.ConversationRequest{NoteID: "asd", Limit: 20},
		},
	)
}

func ExampleService_Conversation() {
	client, _ := misskey.NewClientWithOptions(misskey.WithSimpleConfig("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN")))
	client.LogLevel(logrus.DebugLevel)

	noteList, err := client.Notes().Conversation(notes.ConversationRequest{
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
