package folders_test

import (
	"log"
	"os"

	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/services/drive/folders"
)

func ExampleService_Create() {
	client, _ := misskey.NewClientWithOptions(misskey.WithSimpleConfig("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN")))

	folder, err := client.Drive().Folder().Create(folders.CreateRequest{
		Name: "Test with Go library",
	})
	if err != nil {
		log.Printf("[Drive/Folder/Create] %s", err)

		return
	}

	log.Printf(
		"[Drive/Folder/Create] '%s' folder created. (%s)",
		folder.Name,
		folder.ID,
	)
}
