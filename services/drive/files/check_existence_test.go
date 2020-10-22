package files_test

import (
	"log"
	"os"

	"github.com/yitsushi/go-misskey"
)

func ExampleService_CheckExistence() {
	client := misskey.NewClient("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN"))

	hash := "e960345a4fd3d8413ade5bf1104b1480"

	found, err := client.Drive().File().CheckExistence(hash)
	if err != nil {
		log.Printf("[Drive/File/CheckExistence] %s", err)

		return
	}

	if found {
		log.Printf("[Drive/File/CheckExistence] %s exists.", hash)
	} else {
		log.Printf("[Drive/File/CheckExistence] %s does not exist.", hash)
	}
}
