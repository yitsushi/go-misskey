package main

import (
	"log"

	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/models"
	"github.com/yitsushi/go-misskey/services/antennas"
)

func listAntennas(c *misskey.Client) {
	ants, _ := c.Antennas().List()
	for _, ant := range ants {
		log.Printf("[Antennas/List] <%s> %s", ant.ID, ant.Name)
	}
}

func notesAntennas(c *misskey.Client, antennaID string) {
	notes, err := c.Antennas().Notes(&antennas.NotesOptions{
		AntennaID: antennaID,
	})
	if err != nil {
		log.Printf("[Antennas/Notes] %s", err)
		return
	}

	log.Printf("[Antennas/Notes] %d notes.", len(notes))
}

func createAntenna(c *misskey.Client) models.Antenna {
	antenna, err := c.Antennas().Create(&antennas.CreateOptions{
		Name:            "test",
		Source:          models.AllSrc,
		UserListID:      nil,
		UserGroupID:     nil,
		Keywords:        [][]string{{"update", "what"}, {"stuff"}},
		ExcludeKeywords: [][]string{},
		Users:           []string{},
		CaseSensitive:   false,
		WithReplies:     true,
		WithOnlyFile:    true,
		Notify:          false,
	})
	if err != nil {
		log.Printf("[Antennas/Create] %s", err)
		return models.Antenna{}
	}

	log.Printf("[Antennas/Create] %s created", antenna.Name)

	return antenna
}

func deleteAntenna(c *misskey.Client, id string) {
	_, err := c.Antennas().Delete(id)
	if err == nil {
		log.Printf("[Antennas/Delete] %s deleted, no errors", id)
	} else {
		log.Printf("[Antennas/Delete] Can't delete resource: %s (%s)", err, id)
	}
}

func updateAntenna(c *misskey.Client, antenna *models.Antenna) {
	updated, err := c.Antennas().UpdateAntenna(antenna)
	if err != nil {
		log.Printf("[Antennas/Update] %s", err)
	} else {
		log.Printf("[Antennas/Update] %s updated", updated.Name)
	}
}

func showAntenna(c *misskey.Client, id string) {
	ant, err := c.Antennas().Show(id)
	if err != nil {
		log.Printf("[Antennas/Show] %s", err)
	} else {
		log.Println(ant.Keywords)
	}
}

func antenna(c *misskey.Client) {
	antenna := createAntenna(c)

	listAntennas(c)
	showAntenna(c, antenna.ID)
	notesAntennas(c, antenna.ID)

	antenna.Keywords = append(antenna.Keywords, []string{"thisone"})
	updateAntenna(c, &antenna)
	showAntenna(c, antenna.ID)

	deleteAntenna(c, antenna.ID)
}
