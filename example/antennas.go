package main

import (
	"log"

	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/services/antennas"
)

func antenna(c *misskey.Client) {
	antenna, err := c.Antennas().Create(&antennas.CreateOptions{
		Name:            "test",
		Source:          antennas.AllSrc,
		UserListID:      nil,
		UserGroupID:     nil,
		Keywords:        []string{"update what", "stuff"},
		ExcludeKeywords: []string{},
		Users:           []string{},
		CaseSensitive:   false,
		WithReplies:     true,
		WithOnlyFile:    true,
		Notify:          false,
	})
	if err != nil {
		log.Printf("[Antennas/Create] %s", err)
		return
	}

	log.Printf("[Antennas/Create] %s created", antenna.Name)

	antennas, _ := c.Antennas().List()
	for _, ant := range antennas {
		log.Printf("[Antennas/List] %s", ant.Name)
	}

	_, err = c.Antennas().Delete(antenna.ID)
	if err == nil {
		log.Printf("[Antennas/Delete] %s deleted, no errors", antenna.ID)
	} else {
		log.Printf("[Antennas/Delete] Can't delete resource: %s (%s)", err, antenna.ID)
	}
}
