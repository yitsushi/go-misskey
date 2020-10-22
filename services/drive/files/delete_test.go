package files_test

import (
	"log"
	"os"

	"github.com/yitsushi/go-misskey"
)

func ExampleService_Delete() {
	client := misskey.NewClient("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN"))

	err := client.Drive().File().Delete("8a0snrdwsy")
	if err != nil {
		log.Printf("[Drive/File/Delete] %s", err)

		return
	}
}
