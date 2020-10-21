package main

import (
	"log"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
	"github.com/yitsushi/go-misskey/services/drive/files"
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

	// Fetch changes.
	folder, _ = client.Drive().Folder().Show(folder.ID)

	file, err := driveMutableCreateFile(client, folder)
	if err == nil {
		file.Name = core.NewString("I'm not test-filename anymore")

		_, _ = driveMutableUpdateFile(client, file)

		// Fetch changes.
		file, _ = client.Drive().File().Show(&files.ShowOptions{
			FileID: file.ID,
		})

		driveMutableDeleteFile(client, file)
	}

	driveMutableDeleteFolder(client, folder)
}

func driveMutableCreateFile(c *misskey.Client, folder models.Folder) (models.File, error) {
	file, err := c.Drive().File().CreateFromURL(&files.CreateFromURLOptions{
		Name:     "test-filename",
		FolderID: folder.ID,
		URL:      "https://www.wallpaperup.com/uploads/wallpapers/2014/01/23/235641/862478b1ad52546192af60ff03efbde9-700.jpg", //nolint:lll
	})
	if err != nil {
		log.Printf("[Drive] Error happened: %s", err)

		return models.File{}, err
	}

	log.Printf("[Drive] %s file uploaded. (%s)", core.StringValue(file.Name), file.ID)

	return file, err
}

func driveMutableUpdateFile(c *misskey.Client, file models.File) (models.File, error) {
	file, err := c.Drive().File().Update(file)
	if err != nil {
		log.Printf("[Drive] Error happened: %s", err)
	}

	log.Printf("[Drive] %s file updated.", file.ID)

	return file, err
}

func driveMutableDeleteFile(c *misskey.Client, file models.File) {
	err := c.Drive().File().Delete(file.ID)
	if err != nil {
		log.Printf("[Drive] Error happened: %s", err)

		return
	}

	log.Printf("[Drive] %s file deleted. (%s)", core.StringValue(file.Name), file.ID)
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
