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

func TestService_Renotes(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/notes/renotes",
		RequestData:  &notes.RenotesRequest{},
		ResponseFile: "renotes.json",
		StatusCode:   http.StatusOK,
	})

	renotes, err := client.Notes().Renotes(notes.RenotesRequest{
		Limit:  10,
		NoteID: "asd",
	})
	if !assert.NoError(t, err) {
		return
	}

	assert.Len(t, renotes, 2)

	assert.Equal(t, "8du8xuvc67", renotes[0].ID)
	assert.Equal(t, "83tcps73of", renotes[0].User.ID)
}

func TestRenotesRequest_Validate(t *testing.T) {
	test.ValidateRequests(
		t,
		[]core.BaseRequest{
			notes.RenotesRequest{},
			notes.RenotesRequest{NoteID: "asd"},
		},
		[]core.BaseRequest{
			notes.RenotesRequest{NoteID: "asd", Limit: 20},
		},
	)
}

func ExampleService_Renotes() {
	client, _ := misskey.NewClientWithOptions(misskey.WithSimpleConfig("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN")))

	renotes, err := client.Notes().Renotes(notes.RenotesRequest{
		NoteID: "8dsk7x47y3",
		Limit:  10,
	})
	if err != nil {
		log.Printf("[Notes/Renotes] %s", err)

		return
	}

	for _, renote := range renotes {
		log.Printf("[Notes/Renotes] %s", renote.User.Name)
	}
}
