package main

import (
	"log"

	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/core"
)

func main() {
	client := misskey.NewClient("https://slippy.xyz", "hy9u5ohC4jgykfybi2LDCpm7u3RAVLOE")

	listAnnouncements(client)
	printMeta(client)
	printStats(client)
}

func listAnnouncements(client *misskey.Client) {
	announcements, err := client.Announcements(
		&misskey.AnnouncementOptions{
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
		log.Println(core.StringValue(announcement.Title))
	}
}

func printMeta(client *misskey.Client) {
	meta, err := client.Meta(true)
	if err != nil {
		log.Printf("[Meta] Error happened: %s", err)
		return
	}

	log.Println(core.StringValue(meta.Name))

	for _, emoji := range meta.Emojis {
		log.Printf("[Emoji] %s", core.StringValue(emoji.Name))
	}

	log.Printf("[Feature] Registration:   %s", boolStatusToString(meta.Features.Registration))
	log.Printf("[Feature] LocalTimeLine:  %s", boolStatusToString(meta.Features.LocalTimeLine))
	log.Printf("[Feature] GlobalTimeLine: %s", boolStatusToString(meta.Features.GlobalTimeLine))
	log.Printf("[Feature] Elasticsearch:  %s", boolStatusToString(meta.Features.Elasticsearch))
	log.Printf("[Feature] Hcaptcha:       %s", boolStatusToString(meta.Features.Hcaptcha))
	log.Printf("[Feature] Recaptcha:      %s", boolStatusToString(meta.Features.Recaptcha))
	log.Printf("[Feature] ObjectStorage:  %s", boolStatusToString(meta.Features.ObjectStorage))
	log.Printf("[Feature] Twitter:        %s", boolStatusToString(meta.Features.Twitter))
	log.Printf("[Feature] Github:         %s", boolStatusToString(meta.Features.Github))
	log.Printf("[Feature] Discord:        %s", boolStatusToString(meta.Features.Discord))
	log.Printf("[Feature] ServiceWorker:  %s", boolStatusToString(meta.Features.ServiceWorker))
	log.Printf("[Feature] MiAuth:         %s", boolStatusToString(meta.Features.MiAuth))
}

func printStats(client *misskey.Client) {
	stats, err := client.Stats()
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

func boolStatusToString(v bool) string {
	if v {
		return "enabled"
	}

	return "disabled"
}
