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

func TestService_Reactions(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/notes/reactions",
		RequestData:  &notes.ReactionsRequest{},
		ResponseFile: "reactions.json",
		StatusCode:   http.StatusOK,
	})

	reactions, err := client.Notes().Reactions(notes.ReactionsRequest{
		NoteID: "asd",
		Limit:  20,
	})
	if !assert.NoError(t, err) {
		return
	}

	assert.Len(t, reactions, 1)

	assert.Equal(t, "8dsrstsxbi", reactions[0].ID)
	assert.Equal(t, "syuilo", reactions[0].User.Username)
}

func TestReactionsRequest_Validate(t *testing.T) {
	test.ValidateRequests(
		t,
		[]core.BaseRequest{
			notes.ReactionsRequest{},
			notes.ReactionsRequest{NoteID: "asd"},
		},
		[]core.BaseRequest{
			notes.ReactionsRequest{NoteID: "asd", Limit: 20},
		},
	)
}

func ExampleService_Reactions() {
	client := misskey.NewClient("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN"))

	reactions, err := client.Notes().Reactions(notes.ReactionsRequest{
		NoteID: "8dsk7x47y3",
		Limit:  10,
	})
	if err != nil {
		log.Printf("[Notes/Reactions] %s", err)

		return
	}

	for _, reaction := range reactions {
		log.Printf("[Notes/Reactions] <%s> %s", reaction.User.Username, reaction.Type)
	}
}
