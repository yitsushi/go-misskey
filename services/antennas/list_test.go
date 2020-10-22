package antennas_test

import (
	"log"
	"os"

	"github.com/yitsushi/go-misskey"
)

func ExampleService_List() {
	client := misskey.NewClient("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN"))

	antennas, err := client.Antennas().List()
	if err != nil {
		log.Printf("[Antennas/List] %s", err)

		return
	}

	for _, antenna := range antennas {
		log.Printf("[Antennas/List] %s created", antenna.Name)
	}
}
