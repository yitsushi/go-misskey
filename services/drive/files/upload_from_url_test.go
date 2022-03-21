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

func TestService_UploadFromURL(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/drive/files/upload-from-url",
		RequestData:  &files.UploadFromURLRequest{},
		ResponseFile: "empty",
		StatusCode:   http.StatusNoContent,
	})

	err := client.Drive().File().UploadFromURL(files.UploadFromURLRequest{
		URL: "test",
	})
	assert.NoError(t, err)
}

func ExampleService_UploadFromURL() {
	client, _ := misskey.NewClientWithOptions(misskey.WithSimpleConfig("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN")))

	// Just don't use this one, use CreateFromURL instead.
	err := client.Drive().File().UploadFromURL(files.UploadFromURLRequest{
		URL:         "https://www.wallpaperup.com/uploads/wallpapers/2014/01/23/235641/862478b1ad52546192af60ff03efbde9-700.jpg", //nolint:lll
		Name:        "test-filename",
		IsSensitive: false,
		Force:       false,
	})
	if err != nil {
		log.Printf("[Drive/File/UploadFromURL] %s", err)

		return
	}
}
