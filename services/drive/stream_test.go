package drive_test

import (
	"log"
	"os"

	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/services/drive"
)

func ExampleService_Stream() {
	client, _ := misskey.NewClientWithOptions(misskey.WithSimpleConfig("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN")))

	fileList, err := client.Drive().Stream(drive.StreamRequest{
		Limit: drive.DefaultListLimit,
	})
	if err != nil {
		log.Printf("[Drive/Stream] %s", err)

		return
	}

	log.Printf("[Drive/Stream] Number of files: %d", len(fileList))
}
