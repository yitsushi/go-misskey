package files_test

import (
	"log"
	"os"

	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/services/drive/files"
)

func ExampleService_Show_byID() {
	client, _ := misskey.NewClientWithOptions(misskey.WithSimpleConfig("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN")))

	file, err := client.Drive().File().Show(files.ShowRequest{
		FileID: "8a0snrdwsy",
	})
	if err != nil {
		log.Printf("[Drive/File/Show] %s", err)

		return
	}

	log.Printf(
		"[Drive/File/Show] <%s> %s",
		file.CreatedAt,
		core.StringValue(file.Name),
	)
}

func ExampleService_Show_byURL() {
	client, _ := misskey.NewClientWithOptions(misskey.WithSimpleConfig("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN")))

	file, err := client.Drive().File().Show(files.ShowRequest{
		URL: "https://slippy.xyz/files/7387e4d8-5c44-450d-aa85-9a89a580696e",
	})
	if err != nil {
		log.Printf("[Drive/File/Show] %s", err)

		return
	}

	log.Printf(
		"[Drive/File/Show] <%s> %s",
		file.CreatedAt,
		core.StringValue(file.Name),
	)
}
