package drive_test

import (
	"log"
	"os"

	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/services/drive"
)

func ExampleService_Files() {
	client := misskey.NewClient("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN"))

	fileList, err := client.Drive().Files(drive.FilesRequest{
		Limit: drive.defaultListLimit,
	})
	if err != nil {
		log.Printf("[Drive/Files] %s", err)

		return
	}

	log.Printf("[Drive/Files] Number of files: %d", len(fileList))
}
