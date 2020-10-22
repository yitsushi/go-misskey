package antennas_test

import (
	"log"
	"os"

	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/services/antennas"
)

func ExampleService_Notes() {
	client := misskey.NewClient("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN"))

	notes, err := client.Antennas().Notes(&antennas.NotesOptions{
		AntennaID: "8dbpybhulw",
		Limit:     100,
	})
	if err != nil {
		log.Println(err)

		return
	}

	for _, note := range notes {
		log.Printf(note.Text)
	}
}
