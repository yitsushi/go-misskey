package folders_test

import (
	"log"
	"os"

	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/services/drive/folders"
)

func ExampleService_Find() {
	client := misskey.NewClient("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN"))

	folderList, err := client.Drive().Folder().Find(&folders.FindOptions{
		Name: "Board Games",
	})
	if err != nil {
		log.Printf("[Drive/Folder/Find] %s", err)

		return
	}

	for _, folder := range folderList {
		log.Printf(
			"[Drive/Folder/Find] <%s> %s -> %s",
			folder.CreatedAt,
			folder.ID,
			folder.Name,
		)
	}
}
