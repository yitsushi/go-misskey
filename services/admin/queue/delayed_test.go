package queue_test

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/services/admin/queue"
	"github.com/yitsushi/go-misskey/test"
)

func TestService_InboxDelayed(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/admin/queue/inbox-delayed",
		RequestData:  &queue.StatsRequest{},
		ResponseFile: "inbox-delayed.json",
		StatusCode:   http.StatusOK,
	})

	response, err := client.Admin().Queue().InboxDelayed()

	if !assert.NoError(t, err) {
		return
	}

	assert.Len(t, response, 1)

	assert.Equal(t, "quey.org", response[0].Host)
	assert.Equal(t, int64(2), response[0].Count)
}

func TestService_DeliverDelayed(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/admin/queue/deliver-delayed",
		RequestData:  &queue.StatsRequest{},
		ResponseFile: "deliver-delayed.json",
		StatusCode:   http.StatusOK,
	})

	response, err := client.Admin().Queue().DeliverDelayed()

	if !assert.NoError(t, err) {
		return
	}

	assert.Len(t, response, 1)

	assert.Equal(t, "monads.online", response[0].Host)
	assert.Equal(t, int64(1), response[0].Count)
}

func TestDelayedRequest_Validate(t *testing.T) {
	test.ValidateRequests(
		t,
		[]core.BaseRequest{},
		[]core.BaseRequest{
			queue.DelayedRequest{},
		},
	)
}

func ExampleService_DeliverDelayed() {
	client := misskey.NewClient("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN"))

	response, err := client.Admin().Queue().DeliverDelayed()
	if err != nil {
		log.Printf("[Admin/Queue] %s", err)

		return
	}

	for _, item := range response {
		log.Printf("%s => %d", item.Host, item.Count)
	}
}

func ExampleService_InboxDelayed() {
	client := misskey.NewClient("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN"))

	response, err := client.Admin().Queue().InboxDelayed()
	if err != nil {
		log.Printf("[Admin/Queue] %s", err)

		return
	}

	for _, item := range response {
		log.Printf("%s => %d", item.Host, item.Count)
	}
}
