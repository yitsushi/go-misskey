package favorites_test

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/services/notes/favorites"
	"github.com/yitsushi/go-misskey/test"
)

func TestService_Create(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/notes/favorites/create",
		RequestData:  &favorites.CreateRequest{},
		ResponseFile: "empty",
		StatusCode:   http.StatusNoContent,
	})

	err := client.Notes().Favorites().Create("noteid")
	assert.NoError(t, err)
}

func ExampleService_Create() {
	client := misskey.NewClient("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN"))
	client.LogLevel(logrus.DebugLevel)

	err := client.Notes().Favorites().Create("noteid")
	if err != nil {
		log.Printf("[Notes] Error happened: %s", err)

		return
	}
}
