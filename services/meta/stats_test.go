package meta_test

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/services/meta"
	"github.com/yitsushi/go-misskey/test"
)

func TestService_Stats(t *testing.T) {
	mockClient := test.SimpleMockEndpoint(&test.SimpleMockOptions{
		Endpoint:     "/api/stats",
		RequestData:  &meta.StatsRequest{},
		ResponseFile: "auth/stats.json",
		StatusCode:   http.StatusOK,
	})

	client := misskey.NewClient("https://localhost", "thisistoken")
	client.HTTPClient = mockClient

	response, err := client.Meta().Stats()
	if !assert.NoError(t, err) {
		return
	}

	assert.EqualValues(t, 1990, response.Instances)
	assert.EqualValues(t, 1073029, response.NotesCount)
	assert.EqualValues(t, 1446, response.OriginalNotesCount)
	assert.EqualValues(t, 19588, response.UsersCount)
	assert.EqualValues(t, 59, response.OriginalUsersCount)
	assert.EqualValues(t, 388490580, response.DriveUsageLocal.Bytes())
	assert.EqualValues(t, 16218865643, response.DriveUsageRemote.Bytes())
}

func ExampleService_Stats() {
	client := misskey.NewClient("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN"))

	stats, err := client.Meta().Stats()
	if err != nil {
		log.Printf("[Meta] Error happened: %s", err)

		return
	}

	log.Printf("[Stats] Instances:          %d", stats.Instances)
	log.Printf("[Stats] NotesCount:         %d", stats.NotesCount)
	log.Printf("[Stats] UsersCount:         %d", stats.UsersCount)
	log.Printf("[Stats] OriginalNotesCount: %d", stats.OriginalNotesCount)
	log.Printf("[Stats] OriginalUsersCount: %d", stats.OriginalUsersCount)
}
