package polls_test

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/services/notes/polls"
	"github.com/yitsushi/go-misskey/test"
)

func TestService_Vote(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/notes/polls/vote",
		RequestData:  &polls.VoteRequest{},
		ResponseFile: "empty",
		StatusCode:   http.StatusNoContent,
	})

	err := client.Notes().Polls().Vote("noteid", 5)
	if !assert.NoError(t, err) {
		return
	}
}

func ExampleService_Vote() {
	client := misskey.NewClient("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN"))
	client.LogLevel(logrus.DebugLevel)

	err := client.Notes().Polls().Vote("noteid", 5)
	if err != nil {
		log.Printf("[Notes] Error happened: %s", err)

		return
	}
}
