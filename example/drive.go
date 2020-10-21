package main

import (
	"log"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/services/drive"
	"github.com/yitsushi/go-misskey/services/drive/files"
	"github.com/yitsushi/go-misskey/services/drive/folders"
)

const driveQueryLimit = 3

func driveEndpoints() {
	client := misskey.NewClient("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN"))
	client.LogLevel(logrus.ErrorLevel)

	driveInformation(client)
	driveFolders(client)
	driveFiles(client)
	driveStream(client)
	driveFileAttachedNotes(client)
	driveFileCheckExistence(client)
	driveFileFindByHash(client)
	driveFileFind(client)
	driveFolderFind(client)
	driveFolderShow(client)
	driveFileShow(client)
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
	folderList, err := client.Drive().Folders(&drive.FoldersOptions{
		Limit: driveQueryLimit,
	})
	if err != nil {
		log.Printf("[Drive] Error happened: %s", err)

		return
	}

	for _, folder := range folderList {
		log.Printf("<%s> [%s] %s", folder.ID, core.StringValue(folder.ParentID), folder.Name)
	}
}

func driveFiles(client *misskey.Client) {
	fileList, err := client.Drive().Files(&drive.FilesOptions{
		Limit:    driveQueryLimit,
		FolderID: core.NewString("8dmwq3bhtw"),
	})
	if err != nil {
		log.Printf("[Drive] Error happened: %s", err)

		return
	}

	for _, file := range fileList {
		if file.FolderID != nil {
			log.Printf("<%s> [%s] <%s> %s", file.ID, *file.FolderID, file.Type, *file.Name)
		} else {
			log.Printf("<%s> [    ] <%s> %s", file.ID, file.Type, *file.Name)
		}
	}
}

func driveStream(client *misskey.Client) {
	fileList, err := client.Drive().Stream(&drive.StreamOptions{
		Limit: driveQueryLimit,
		Type:  "image/gif",
	})
	if err != nil {
		log.Printf("[Drive] Error happened: %s", err)

		return
	}

	for _, file := range fileList {
		log.Printf("<%s> [%s] <%s> %s", file.ID, core.StringValue(file.FolderID), file.Type, *file.Name)
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

func driveFileCheckExistence(c *misskey.Client) {
	checkFound := "e960345a4fd3d8413ade5bf1104b1480"
	checkNotFound := "ffffffffffffffffffffffffffffffff"

	found, err := c.Drive().File().CheckExistence(checkFound)
	if err != nil {
		log.Printf("[Drive] Error happened: %s", err)

		return
	}

	notFound, err := c.Drive().File().CheckExistence(checkNotFound)
	if err != nil {
		log.Printf("[Drive] Error happened: %s", err)

		return
	}

	log.Println(found)
	log.Println(notFound)
}

func driveFileFindByHash(c *misskey.Client) {
	check := "e960345a4fd3d8413ade5bf1104b1480"

	fileList, err := c.Drive().File().FindByHash(check)
	if err != nil {
		log.Printf("[Drive] Error happened: %s", err)

		return
	}

	for _, file := range fileList {
		log.Printf("Filename (find by hash): %s/%s", *file.FolderID, *file.Name)
	}
}

func driveFileFind(c *misskey.Client) {
	fileList, err := c.Drive().File().Find(&files.FindOptions{
		Name:     "IMG_20200722_123302.jpg",
		FolderID: core.NewString("8dmwq3bhtw"),
	})
	if err != nil {
		log.Printf("[Drive] Error happened: %s", err)

		return
	}

	for _, file := range fileList {
		log.Printf("Filename (find by name): %s/%s", *file.FolderID, *file.Name)
	}
}

func driveFolderFind(c *misskey.Client) {
	folderList, err := c.Drive().Folder().Find(&folders.FindOptions{
		Name: "Board Games",
	})
	if err != nil {
		log.Printf("[Drive] Error happened: %s", err)

		return
	}

	for _, folder := range folderList {
		log.Printf(
			"<%s> [%s] %s",
			folder.CreatedAt,
			folder.ID,
			folder.Name,
		)
	}
}

func driveFolderShow(c *misskey.Client) {
	folder, err := c.Drive().Folder().Show("8dmwisynnu")
	if err != nil {
		log.Printf("[Drive] Error happened: %s", err)

		return
	}

	log.Printf(
		"<%s> %s (%d files and %d folders)",
		folder.CreatedAt,
		folder.Name,
		folder.FilesCount,
		folder.FoldersCount,
	)
}

func driveFileShow(c *misskey.Client) {
	driveFileShowByID(c)
	driveFileShowByURL(c)
}

func driveFileShowByID(c *misskey.Client) {
	file, err := c.Drive().File().Show(&files.ShowOptions{
		FileID: "8a0snrdwsy",
	})
	if err != nil {
		log.Printf("[Drive] Error happened: %s", err)

		return
	}

	log.Printf(
		"<%s> %s",
		file.CreatedAt,
		*file.Name,
	)
}

func driveFileShowByURL(c *misskey.Client) {
	file, err := c.Drive().File().Show(&files.ShowOptions{
		URL: "https://slippy.xyz/files/21fe9ecf-20f1-40e6-a92c-11b9d9072d88",
	})
	if err != nil {
		log.Printf("[Drive] Error happened: %s", err)

		return
	}

	log.Printf(
		"<%s> %s",
		file.CreatedAt,
		*file.Name,
	)
}
