package files_test

import (
	"log"
	"os"

	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/services/drive/files"
)

func ExampleService_CreateFromURL() {
	client := misskey.NewClient("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN"))

	file, err := client.Drive().File().CreateFromURL(&files.CreateFromURLOptions{
		Name:     "test-filename",
		FolderID: "8dmwisynnu",
		URL:      "https://www.wallpaperup.com/uploads/wallpapers/2014/01/23/235641/862478b1ad52546192af60ff03efbde9-700.jpg", //nolint:lll
	})
	if err != nil {
		log.Printf("[Drive/File/CreateFromURL] %s", err)

		return
	}

	log.Printf("[Drive/File/CreateFromURL] %s uploaded.", core.StringValue(file.Name))
}
