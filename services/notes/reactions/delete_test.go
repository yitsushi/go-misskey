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

func TestService_Delete(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/notes/reactions/delete",
		RequestData:  &reactions.DeleteRequest{},
		ResponseFile: "empty",
		StatusCode:   http.StatusNoContent,
	})

	err := client.Notes().Reactions().Delete("asd")
	if !assert.NoError(t, err) {
		return
	}
}

func TestDeleteRequest_Validate(t *testing.T) {
	test.ValidateRequests(
		t,
		[]core.BaseRequest{
			reactions.DeleteRequest{},
		},
		[]core.BaseRequest{
			reactions.DeleteRequest{NoteID: "asd"},
		},
	)
}

func ExampleService_Delete() {
	client, _ := misskey.NewClientWithOptions(misskey.WithSimpleConfig("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN")))

	err := client.Notes().Reactions().Delete("8dsk7x47y3")
	if err != nil {
		log.Printf("[Notes/Reaction/Delete] %s", err)

		return
	}
}
