package meta_test

import (
	"log"
	"os"

	"github.com/yitsushi/go-misskey"
)

func ExampleService_Stats() {
	client := misskey.NewClient("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN"))

	stats, err := client.Meta().Stats()
	if err != nil {
		log.Printf("[Meta] Error happened: %s", err)

		return
	}

	log.Printf("[Stats] Instances:          %d", stats.Instances)
	log.Printf("[Stats] NotesCount:         %d", stats.NotesCount)
	log.Printf("[Stats] UsersCount:         %d", stats.UsersCount)
	log.Printf("[Stats] OriginalNotesCount: %d", stats.OriginalNotesCount)
	log.Printf("[Stats] OriginalUsersCount: %d", stats.OriginalUsersCount)
}
