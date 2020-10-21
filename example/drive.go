package main

import (
	"log"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/services/drive"
)

const driveQueryLimit = 3

func driveEndpoints() {
	client := misskey.NewClient("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN"))
	client.LogLevel(logrus.ErrorLevel)

	driveInformation(client)
	driveFolders(client)
	driveFiles(client)
	driveFileAttachedNotes(client)
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

func driveFolders(client *misskey.Client) {
	folders, err := client.Drive().Folders(&drive.FoldersOptions{
		Limit: driveQueryLimit,
	})
	if err != nil {
		log.Printf("[Drive] Error happened: %s", err)
		return
	}

	for _, folder := range folders {
		log.Printf("<%s> [%s] %s", folder.ID, core.StringValue(folder.ParentID), folder.Name)
	}
}

func driveFiles(client *misskey.Client) {
	files, err := client.Drive().Files(&drive.FilesOptions{
		Limit:    driveQueryLimit,
		FolderID: core.NewString("8dmwq3bhtw"),
	})
	if err != nil {
		log.Printf("[Drive] Error happened: %s", err)
		return
	}

	for _, file := range files {
		if file.FolderID != nil {
			log.Printf("<%s> [%s] <%s> %s", file.ID, *file.FolderID, file.Type, *file.Name)
		} else {
			log.Printf("<%s> [    ] <%s> %s", file.ID, file.Type, *file.Name)
		}
	}
}

func driveFileAttachedNotes(c *misskey.Client) {
	notes, err := c.Drive().File().AttachedNotes("8a0snrdwsy")
	if err != nil {
		log.Printf("[Drive] Error happened: %s", err)
		return
	}

	for _, note := range notes {
		log.Printf("[%s] <%s> %s", note.CreatedAt, note.User.Name, note.Text)
	}
}
