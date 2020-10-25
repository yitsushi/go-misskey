package files_test

import (
	"log"
	"os"

	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/services/drive/files"
)

func ExampleService_Update() {
	client := misskey.NewClient("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN"))

	file, err := client.Drive().File().Show(files.ShowRequest{
		FileID: "8a0snrdwsy",
	})
	if err != nil {
		log.Printf("[Drive/File/Update] %s", err)

		return
	}

	file.FolderID = core.NewString("8dmwq3bhtw")

	_, err = client.Drive().File().Update(file)
	if err != nil {
		log.Printf("[Drive/File/Update] %s", err)

		return
	}
}
