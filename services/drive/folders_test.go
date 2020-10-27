package drive_test

import (
	"log"
	"os"

	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/services/drive"
)

func ExampleService_Folders() {
	client := misskey.NewClient("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN"))

	folderList, err := client.Drive().Folders(drive.FoldersRequest{
		Limit: drive.defaultListLimit,
	})
	if err != nil {
		log.Printf("[Drive/Folders] %s", err)

		return
	}

	log.Printf("[Drive/Folders] Number of files: %d", len(folderList))
}
