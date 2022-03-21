package antennas_test

import (
	"log"
	"os"

	"github.com/yitsushi/go-misskey"
)

func ExampleService_Update() {
	client, _ := misskey.NewClientWithOptions(misskey.WithSimpleConfig("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN")))

	antenna, err := client.Antennas().Show("8dbpybhulw")
	if err != nil {
		log.Println(err)

		return
	}

	antenna.Keywords = append(antenna.Keywords, []string{"addition"})

	_, err = client.Antennas().UpdateAntenna(&antenna)
	if err != nil {
		log.Println(err)

		return
	}
}
