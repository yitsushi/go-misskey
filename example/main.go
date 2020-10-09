package main

import (
	"os"

	"github.com/yitsushi/go-misskey"
)

func main() {
	client := misskey.NewClient("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN"))

	// listAnnouncements(client)
	// printMeta(client)
	printStats(client)
	// createAntenna(client)
}

func boolStatusToString(v bool) string {
	if v {
		return "enabled"
	}

	return "disabled"
}
