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

func TestService_Stats(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/admin/queue/stats",
		RequestData:  &queue.StatsRequest{},
		ResponseFile: "stats.json",
		StatusCode:   http.StatusOK,
	})

	response, err := client.Admin().Queue().Stats()

	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, int64(2), response.Deliver.Active)
	assert.Equal(t, int64(6), response.Deliver.Paused)
}

func TestStatsRequest_Validate(t *testing.T) {
	test.ValidateRequests(
		t,
		[]core.BaseRequest{},
		[]core.BaseRequest{
			queue.StatsRequest{},
		},
	)
}

func ExampleService_Stats() {
	client, _ := misskey.NewClientWithOptions(misskey.WithSimpleConfig("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN")))

	response, err := client.Admin().Queue().Stats()
	if err != nil {
		log.Printf("[Admin/Queue] %s", err)

		return
	}

	log.Printf("Deliver :: Waiting = %d", response.Deliver.Waiting)
}
