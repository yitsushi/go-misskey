package folders_test

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/services/drive/folders"
	"github.com/yitsushi/go-misskey/test"
)

func TestService_Delete(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/drive/folders/delete",
		RequestData:  &folders.DeleteRequest{},
		ResponseFile: "empty",
		StatusCode:   http.StatusNoContent,
	})

	err := client.Drive().Folder().Delete("test")
	assert.NoError(t, err)
}

func ExampleService_Delete() {
	client := misskey.NewClient("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN"))

	err := client.Drive().Folder().Delete("8dmwq3bhtw")
	if err != nil {
		log.Printf("[Drive/Folder/Delete] %s", err)

		return
	}
}
