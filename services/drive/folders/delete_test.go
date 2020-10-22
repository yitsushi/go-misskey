package folders_test

import (
	"log"
	"os"

	"github.com/yitsushi/go-misskey"
)

func ExampleService_Delete() {
	client := misskey.NewClient("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN"))

	err := client.Drive().Folder().Delete("8dmwq3bhtw")
	if err != nil {
		log.Printf("[Drive/Folder/Delete] %s", err)

		return
	}
}
