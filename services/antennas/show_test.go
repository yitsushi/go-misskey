package antennas_test

import (
	"log"
	"os"

	"github.com/yitsushi/go-misskey"
)

func ExampleService_Show() {
	client := misskey.NewClient("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN"))

	antenna, err := client.Antennas().Show("8dbpybhulw")
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(antenna.Name)
}
