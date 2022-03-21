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

func TestService_Show(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/notes/show",
		RequestData:  &notes.ShowRequest{},
		ResponseFile: "show.json",
		StatusCode:   http.StatusOK,
	})

	note, err := client.Notes().Show("8dsk7x47y3")
	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, "8dsk7x47y3", note.ID)
	assert.Equal(t, "efertone", note.User.Username)
}

func ExampleService_Show() {
	client, _ := misskey.NewClientWithOptions(misskey.WithSimpleConfig("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN")))

	note, err := client.Notes().Show("8dsk7x47y3")
	if err != nil {
		log.Printf("[Notes/Show] %s", err)

		return
	}

	log.Printf("[Notes/Show] <%s> %s", note.User.Username, note.Text)
}
