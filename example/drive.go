package main

import (
	"log"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/services/drive"
)

func driveEndpoints() {
	client := misskey.NewClient("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN"))
	client.LogLevel(logrus.DebugLevel)

	driveInformation(client)
	driveFiles(client)
}

func driveInformation(client *misskey.Client) {
	information, err := client.Drive().Information()
	if err != nil {
		log.Printf("[Drive] Error happened: %s", err)
		return
	}

	log.Printf("Capacity: %.2f MB", information.Capacity.Megabytes())
	log.Printf("Usage:    %.2f MB", information.Usage.Megabytes())
}

func driveFiles(client *misskey.Client) {
	files, err := client.Drive().Files(&drive.FilesOptions{
		Limit: 100,
	})
	if err != nil {
		log.Printf("[Drive] Error happened: %s", err)
		return
	}

	for _, file := range files {
		if file.FolderID != nil {
			log.Printf("[%s] <%s> %s", *file.FolderID, file.Type, *file.Name)
		} else {
			log.Printf("[    ] <%s> %s", file.Type, *file.Name)
		}
	}
}
