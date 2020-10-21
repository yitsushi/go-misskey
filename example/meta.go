package main

import (
	"log"
	"os"
	"strings"

	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/services/meta"
)

func announcements() {
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

func instanceMeta() {
	client := misskey.NewClient("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN"))

	meta, err := client.Meta().InstanceMeta(true)
	if err != nil {
		log.Printf("[Meta] Error happened: %s", err)
		return
	}

	log.Printf("[InstanceMeta/Name] %s", core.StringValue(meta.Name))

	emojiList := []string{}
	for _, emoji := range meta.Emojis {
		emojiList = append(emojiList, core.StringValue(emoji.Name))
	}

	log.Printf("[InstanceMeta/Emoji] %s", strings.Join(emojiList, ", "))

	log.Printf("[InstanceMeta/Feature] Registration:   %s", boolStatusToString(meta.Features.Registration))
	log.Printf("[InstanceMeta/Feature] LocalTimeLine:  %s", boolStatusToString(meta.Features.LocalTimeLine))
	log.Printf("[InstanceMeta/Feature] GlobalTimeLine: %s", boolStatusToString(meta.Features.GlobalTimeLine))
	log.Printf("[InstanceMeta/Feature] Elasticsearch:  %s", boolStatusToString(meta.Features.Elasticsearch))
	log.Printf("[InstanceMeta/Feature] Hcaptcha:       %s", boolStatusToString(meta.Features.Hcaptcha))
	log.Printf("[InstanceMeta/Feature] Recaptcha:      %s", boolStatusToString(meta.Features.Recaptcha))
	log.Printf("[InstanceMeta/Feature] ObjectStorage:  %s", boolStatusToString(meta.Features.ObjectStorage))
	log.Printf("[InstanceMeta/Feature] Twitter:        %s", boolStatusToString(meta.Features.Twitter))
	log.Printf("[InstanceMeta/Feature] Github:         %s", boolStatusToString(meta.Features.Github))
	log.Printf("[InstanceMeta/Feature] Discord:        %s", boolStatusToString(meta.Features.Discord))
	log.Printf("[InstanceMeta/Feature] ServiceWorker:  %s", boolStatusToString(meta.Features.ServiceWorker))
	log.Printf("[InstanceMeta/Feature] MiAuth:         %s", boolStatusToString(meta.Features.MiAuth))
}

func stats() {
	client := misskey.NewClient("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN"))

	stats, err := client.Meta().Stats()
	if err != nil {
		log.Printf("[Meta] Error happened: %s", err)
		return
	}

	log.Printf("[Stats] Instances:          %d", stats.Instances)
	log.Printf("[Stats] NotesCount:         %d", stats.NotesCount)
	log.Printf("[Stats] UsersCount:         %d", stats.UsersCount)
	log.Printf("[Stats] OriginalNotesCount: %d", stats.OriginalNotesCount)
	log.Printf("[Stats] OriginalUsersCount: %d", stats.OriginalUsersCount)
}
