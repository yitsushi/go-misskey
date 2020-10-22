package antennas_test

import (
	"log"
	"os"

	"github.com/yitsushi/go-misskey"
)

func ExampleService_Delete() {
	client := misskey.NewClient("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN"))

	err := client.Antennas().Delete("antenna-id")
	if err != nil {
		log.Printf("[Antennas/Delete] %s", err)

		return
	}

	log.Println("[Antennas/Delete] Done without errors")
}
