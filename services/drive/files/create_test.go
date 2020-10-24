package files_test

import (
	"log"
	"os"

	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/services/drive/files"
)

func ExampleService_Create() {
	client := misskey.NewClient("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN"))

	fileContent := []byte{}

	file, err := client.Drive().File().Create(files.CreateRequest{
		FolderID:    "",
		Name:        "this is the name",
		IsSensitive: false,
		Force:       false,
		Content:     fileContent,
	})
	if err != nil {
		log.Printf("[Drive/File/Create] %s", err)

		return
	}

	log.Printf(
		"[Drive/File/Create] %s file uploaded. (%s)",
		core.StringValue(file.Name),
		file.ID,
	)
}
