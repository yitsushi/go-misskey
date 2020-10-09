package main

import (
	"log"

	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/services/antennas"
)

func createAntenna(c *misskey.Client) {
	resp, err := c.Antennas().Create(&antennas.CreateOptions{
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

	log.Printf("[Antennas/Create] %s created", resp.Name)
}
