package files_test

import (
	"log"
	"os"

	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/services/drive/files"
)

func ExampleService_Find() {
	client := misskey.NewClient("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN"))

	fileList, err := client.Drive().File().Find(&files.FindOptions{
		Name:     "file-i-really-really-want.png",
		FolderID: core.NewString("8dmwq3bhtw"),
	})
	if err != nil {
		log.Printf("[Drive/File/Find] %s", err)

		return
	}

	for _, file := range fileList {
		log.Printf(
			"[Drive/File/Find] %s -> %s",
			core.StringValue(file.FolderID),
			core.StringValue(file.Name),
		)
	}
}
