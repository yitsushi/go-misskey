package files_test

import (
	"log"
	"os"

	"github.com/yitsushi/go-misskey"
)

func ExampleService_AttachedNotes() {
	client, _ := misskey.NewClientWithOptions(misskey.WithSimpleConfig("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN")))

	notes, err := client.Drive().File().AttachedNotes("8a0snrdwsy")
	if err != nil {
		log.Printf("[Drive/File/AttachedNotes] %s", err)

		return
	}

	for _, note := range notes {
		log.Printf("[Drive/File/AttachedNotes] <%s> %s", note.User.Name, note.Text)
	}
}
