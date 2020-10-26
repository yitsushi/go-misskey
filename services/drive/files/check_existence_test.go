package files_test

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/services/drive/files"
	"github.com/yitsushi/go-misskey/test"
)

func TestService_CheckExistence(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/drive/files/check-existence",
		RequestData:  &files.CheckExistenceRequest{},
		ResponseFile: "boolean",
		StatusCode:   http.StatusOK,
	})

	found, err := client.Drive().File().CheckExistence("test")
	if !assert.NoError(t, err) {
		return
	}

	assert.True(t, found)
}

func ExampleService_CheckExistence() {
	client := misskey.NewClient("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN"))

	hash := "e960345a4fd3d8413ade5bf1104b1480"

	found, err := client.Drive().File().CheckExistence(hash)
	if err != nil {
		log.Printf("[Drive/File/CheckExistence] %s", err)

		return
	}

	if found {
		log.Printf("[Drive/File/CheckExistence] %s exists.", hash)
	} else {
		log.Printf("[Drive/File/CheckExistence] %s does not exist.", hash)
	}
}
