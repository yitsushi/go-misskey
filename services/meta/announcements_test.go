package meta_test

import (
	"log"
	"os"

	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/services/meta"
)

func ExampleService_Announcements() {
	client := misskey.NewClient("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN"))

	announcements, err := client.Meta().Announcements(
		&meta.AnnouncementOptions{
			WithUnreads: true,
			SinceID:     "",
			UntilID:     "",
		},
	)
	if err != nil {
		log.Printf("[Announcements] Error happened: %s", err)
		return
	}

	for _, announcement := range announcements {
		log.Printf("[Announcements] %s", core.StringValue(announcement.Title))
	}
}
