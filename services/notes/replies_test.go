package notes_test

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/services/notes"
	"github.com/yitsushi/go-misskey/test"
)

func TestService_Replies(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/notes/replies",
		RequestData:  &notes.RepliesRequest{},
		ResponseFile: "replies.json",
		StatusCode:   http.StatusOK,
	})

	replies, err := client.Notes().Replies(notes.RepliesRequest{
		NoteID: "noteid",
		Limit:  10,
	})
	if !assert.NoError(t, err) {
		return
	}

	assert.Len(t, replies, 1)
	assert.Equal(t, "syuilo", replies[0].User.Username)
}

func TestRepliesRequest_Validate(t *testing.T) {
	test.ValidateRequests(
		t,
		[]core.BaseRequest{
			notes.RepliesRequest{},
			notes.RepliesRequest{NoteID: "asd"},
		},
		[]core.BaseRequest{
			notes.RenotesRequest{NoteID: "asd", Limit: 20},
		},
	)
}

func ExampleService_Replies() {
	client := misskey.NewClient("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN"))

	replies, err := client.Notes().Replies(notes.RepliesRequest{
		NoteID: "8dsk7x47y3",
		Limit:  10,
	})
	if err != nil {
		log.Printf("[Notes/Replies] %s", err)

		return
	}

	for _, reply := range replies {
		log.Printf("[Notes/Replies] %s", reply.User.Name)
	}
}
