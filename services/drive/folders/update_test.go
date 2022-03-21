package folders_test

import (
	"log"
	"os"

	"github.com/yitsushi/go-misskey"
)

func ExampleService_Update() {
	client, _ := misskey.NewClientWithOptions(misskey.WithSimpleConfig("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN")))

	folder, err := client.Drive().Folder().Show("8dmwisynnu")
	if err != nil {
		log.Printf("[Drive/Folder/Update] %s", err)

		return
	}

	folder.Name = "New Name"

	folder, err = client.Drive().Folder().Update(folder)
	if err != nil {
		log.Printf("[Drive/Folder/Update] %s", err)

		return
	}
}
