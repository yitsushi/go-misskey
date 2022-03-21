package folders_test

import (
	"log"
	"os"

	"github.com/yitsushi/go-misskey"
)

func ExampleService_Show() {
	client, _ := misskey.NewClientWithOptions(misskey.WithSimpleConfig("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN")))

	folder, err := client.Drive().Folder().Show("8dmwisynnu")
	if err != nil {
		log.Printf("[Drive/Folder/Show] %s", err)

		return
	}

	log.Printf(
		"[Drive/Folder/Show] <%s> %s (%d files and %d folders)",
		folder.CreatedAt,
		folder.Name,
		folder.FilesCount,
		folder.FoldersCount,
	)
}
