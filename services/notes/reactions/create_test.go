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

func TestService_Create(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/notes/reactions/create",
		RequestData:  &reactions.CreateRequest{},
		ResponseFile: "empty",
		StatusCode:   http.StatusNoContent,
	})

	err := client.Notes().Reactions().Create(reactions.CreateRequest{
		NoteID:   "asd",
		Reaction: "👋",
	})
	if !assert.NoError(t, err) {
		return
	}
}

func TestCreateRequest_Validate(t *testing.T) {
	test.ValidateRequests(
		t,
		[]core.BaseRequest{
			reactions.CreateRequest{},
			reactions.CreateRequest{NoteID: "asd"},
		},
		[]core.BaseRequest{
			reactions.CreateRequest{NoteID: "asd", Reaction: "👋"},
		},
	)
}

func ExampleService_Create() {
	client, _ := misskey.NewClientWithOptions(misskey.WithSimpleConfig("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN")))

	err := client.Notes().Reactions().Create(reactions.CreateRequest{
		NoteID:   "8dsk7x47y3",
		Reaction: "👋",
	})
	if err != nil {
		log.Printf("[Notes/Reaction/Create] %s", err)

		return
	}
}
