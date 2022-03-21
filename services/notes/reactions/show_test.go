package reactions_test

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/services/notes/reactions"
	"github.com/yitsushi/go-misskey/test"
)

func TestService_Show(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/notes/reactions",
		RequestData:  &reactions.ShowRequest{},
		ResponseFile: "show.json",
		StatusCode:   http.StatusOK,
	})

	reactions, err := client.Notes().Reactions().Show(reactions.ShowRequest{
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

func TestShowRequest_Validate(t *testing.T) {
	test.ValidateRequests(
		t,
		[]core.BaseRequest{
			reactions.ShowRequest{},
			reactions.ShowRequest{NoteID: "asd"},
		},
		[]core.BaseRequest{
			reactions.ShowRequest{NoteID: "asd", Limit: 20},
		},
	)
}

func ExampleService_Show() {
	client, _ := misskey.NewClientWithOptions(misskey.WithSimpleConfig("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN")))

	reactions, err := client.Notes().Reactions().Show(reactions.ShowRequest{
		NoteID: "8dsk7x47y3",
		Limit:  10,
	})
	if err != nil {
		log.Printf("[Notes/Show] %s", err)

		return
	}

	for _, reaction := range reactions {
		log.Printf("[Notes/Show] <%s> %s", reaction.User.Username, reaction.Type)
	}
}
