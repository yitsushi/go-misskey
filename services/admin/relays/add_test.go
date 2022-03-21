package relays_test

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
	"github.com/yitsushi/go-misskey/services/admin/relays"
	"github.com/yitsushi/go-misskey/test"
)

func TestService_Add(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/admin/relays/add",
		RequestData:  &relays.AddRequest{},
		ResponseFile: "add.json",
		StatusCode:   http.StatusOK,
	})

	relay, err := client.Admin().Relays().Add("https://something.tld/inbox")

	assert.NoError(t, err)
	assert.Equal(t, models.RelayStatusRequesting, relay.Status)
}

func TestAddRequest_Validate(t *testing.T) {
	test.ValidateRequests(
		t,
		[]core.BaseRequest{
			relays.AddRequest{},
		},
		[]core.BaseRequest{
			relays.AddRequest{Inbox: "https://something.tld/inbox"},
		},
	)
}

func ExampleService_Add() {
	client, _ := misskey.NewClientWithOptions(misskey.WithSimpleConfig("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN")))

	relay, err := client.Admin().Relays().Add("https://something.tld/inbox")
	if err != nil {
		log.Printf("[Admin/Relays/Add] %s", err)

		return
	}

	log.Printf("[Admin/Relays/Add] %s", relay.Status)
}
