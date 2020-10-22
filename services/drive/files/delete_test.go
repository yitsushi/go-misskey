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

func TestService_Delete(t *testing.T) {
	mockClient := test.SimpleMockEndpoint(&test.SimpleMockOptions{
		Endpoint:     "/api/drive/files/delete",
		RequestData:  &files.DeleteRequest{},
		ResponseFile: "empty",
		StatusCode:   http.StatusNoContent,
	})

	client := misskey.NewClient("https://localhost", "thisistoken")
	client.HTTPClient = mockClient

	err := client.Drive().File().Delete("test")
	assert.NoError(t, err)
}

func ExampleService_Delete() {
	client := misskey.NewClient("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN"))

	err := client.Drive().File().Delete("8a0snrdwsy")
	if err != nil {
		log.Printf("[Drive/File/Delete] %s", err)

		return
	}
}
