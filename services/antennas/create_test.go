package antennas_test

import (
	"log"
	"os"

	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/models"
	"github.com/yitsushi/go-misskey/services/antennas"
)

func ExampleService_Create() {
	client := misskey.NewClient("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN"))

	resp, err := client.Antennas().Create(antennas.CreateRequest{
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

		return
	}

	log.Printf("[Antennas/Create] %s created", resp.Name)
}
