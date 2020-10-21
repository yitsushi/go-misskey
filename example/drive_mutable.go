package main

import (
	"log"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/models"
	"github.com/yitsushi/go-misskey/services/drive/folders"
)

func driveMutableFlow() {
	client := misskey.NewClient("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN"))
	client.LogLevel(logrus.ErrorLevel)

	folder := driveMutableCreateFolder(client)
	if folder == nil {
		log.Println("[Drive] Abort...")

		return
	}

	driveMutableDeleteFolder(client, folder)
}

func driveMutableCreateFolder(c *misskey.Client) *models.Folder {
	folder, err := c.Drive().Folder().Create(&folders.CreateOptions{
		Name: "Test with Go library",
	})
	if err != nil {
		log.Printf("[Drive] Error happened: %s", err)

		return nil
	}

	log.Printf("[Drive] %s folder created. (%s)", folder.Name, folder.ID)

	return &folder
}

func driveMutableDeleteFolder(c *misskey.Client, folder *models.Folder) {
	err := c.Drive().Folder().Delete(folder.ID)
	if err != nil {
		log.Printf("[Drive] Error happened: %s", err)

		return
	}

	log.Printf("[Drive] %s folder deleted. (%s)", folder.Name, folder.ID)
}
