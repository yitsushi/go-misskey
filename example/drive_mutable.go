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

	folder, err := driveMutableCreateFolder(client)
	if err != nil {
		log.Println("[Drive] Abort...")

		return
	}

	folder.Name = "[Updated] Test with Go library"

	_, _ = driveMutableUpdateFolder(client, folder)

	folder, _ = client.Drive().Folder().Show(folder.ID)

	driveMutableDeleteFolder(client, folder)
}

func driveMutableCreateFolder(c *misskey.Client) (models.Folder, error) {
	folder, err := c.Drive().Folder().Create(&folders.CreateOptions{
		Name: "Test with Go library",
	})
	if err != nil {
		log.Printf("[Drive] Error happened: %s", err)

		return models.Folder{}, err
	}

	log.Printf("[Drive] %s folder created. (%s)", folder.Name, folder.ID)

	return folder, nil
}

func driveMutableDeleteFolder(c *misskey.Client, folder models.Folder) {
	err := c.Drive().Folder().Delete(folder.ID)
	if err != nil {
		log.Printf("[Drive] Error happened: %s", err)

		return
	}

	log.Printf("[Drive] %s folder deleted. (%s)", folder.Name, folder.ID)
}

func driveMutableUpdateFolder(c *misskey.Client, folder models.Folder) (models.Folder, error) {
	folder, err := c.Drive().Folder().Update(folder)
	if err != nil {
		log.Printf("[Drive] Error happened: %s", err)
	}

	log.Printf("[Drive] %s folder updated.", folder.ID)

	return folder, err
}
