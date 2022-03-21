package notes_test

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/services/notes"
	"github.com/yitsushi/go-misskey/test"
)

func TestService_State(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/notes/state",
		RequestData:  &notes.StateRequest{},
		ResponseFile: "state.json",
		StatusCode:   http.StatusOK,
	})

	state, err := client.Notes().State("8dsk7x47y3")
	if !assert.NoError(t, err) {
		return
	}

	assert.True(t, state.IsFavorited)
	assert.False(t, state.IsWatching)
}

func ExampleService_State() {
	client, _ := misskey.NewClientWithOptions(misskey.WithSimpleConfig("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN")))

	state, err := client.Notes().State("8dsk7x47y3")
	if err != nil {
		log.Printf("[Notes/State] %s", err)

		return
	}

	log.Printf("[Notes/State] %v", state)
}
