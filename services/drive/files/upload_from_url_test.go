package files_test

import (
	"log"
	"os"

	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/services/drive/files"
)

func ExampleService_UploadFromURL() {
	client := misskey.NewClient("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN"))

	// Just don't use this one, use CreateFromURL instead.
	err := client.Drive().File().UploadFromURL(&files.UploadFromURLOptions{
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
