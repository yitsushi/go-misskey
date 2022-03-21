package files_test

import (
	"log"
	"os"

	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/core"
)

func ExampleService_FindByHash() {
	client, _ := misskey.NewClientWithOptions(misskey.WithSimpleConfig("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN")))

	hash := "e960345a4fd3d8413ade5bf1104b1480"

	fileList, err := client.Drive().File().FindByHash(hash)
	if err != nil {
		log.Printf("[Drive/File/FindByHash] %s", err)

		return
	}

	for _, file := range fileList {
		log.Printf(
			"[Drive/File/FindByHash] %s -> %s",
			core.StringValue(file.FolderID),
			core.StringValue(file.Name),
		)
	}
}
