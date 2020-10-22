package meta_test

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/services/meta"
	"github.com/yitsushi/go-misskey/test"
)

func TestService_InstanceMeta(t *testing.T) {
	mockClient := test.SimpleMockEndpoint(
		"/api/meta",
		&meta.InstanceMetaRequest{},
		"anon/instance_meta.json",
	)

	client := misskey.NewClient("https://localhost", "thisistoken")
	client.HTTPClient = mockClient

	response, err := client.Meta().InstanceMeta(false)
	if !assert.NoError(t, err) {
		return
	}

	if !assert.Len(t, response.Emojis, 36) {
		return
	}

	assert.False(t, response.Features.Registration)
}

func TestService_InstanceMeta_detailed(t *testing.T) {
	mockClient := test.SimpleMockEndpoint(
		"/api/meta",
		&meta.InstanceMetaRequest{},
		"anon/instance_meta_detailed.json",
	)

	client := misskey.NewClient("https://localhost", "thisistoken")
	client.HTTPClient = mockClient

	response, err := client.Meta().InstanceMeta(true)
	if !assert.NoError(t, err) {
		return
	}

	if !assert.Len(t, response.Emojis, 36) {
		return
	}

	assert.True(t, response.Features.Registration)
}

func ExampleService_InstanceMeta() {
	client := misskey.NewClient("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN"))

	meta, err := client.Meta().InstanceMeta(true)
	if err != nil {
		log.Printf("[Meta] Error happened: %s", err)

		return
	}

	log.Printf("[InstanceMeta/Name] %s", core.StringValue(meta.Name))

	for _, emoji := range meta.Emojis {
		log.Printf("[InstanceMeta/Emoji] %s", core.StringValue(emoji.Name))
	}

	boolStatusToString := func(v bool) string {
		if v {
			return "enabled"
		}

		return "disabled"
	}

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
