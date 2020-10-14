package main

import (
	"os"

	"github.com/yitsushi/go-misskey"
)

func main() {
	client := misskey.NewClient("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN"))

	announcements(client)
	instanceMeta(client)
	stats(client)

	antenna(client)
}

func boolStatusToString(v bool) string {
	if v {
		return "enabled"
	}

	return "disabled"
}
