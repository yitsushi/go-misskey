package instance_test

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/services/admin/instance"
	"github.com/yitsushi/go-misskey/test"
)

func TestService_GetTableStats(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/admin/get-table-stats",
		RequestData:  &instance.GetTableStatsRequest{},
		ResponseFile: "get-table-stats.json",
		StatusCode:   http.StatusOK,
	})

	response, err := client.Admin().Instance().GetTableStats()

	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, 66, response.Migrations.Count)
	assert.Equal(t, 32768, response.Migrations.Size)
}

// ExampleService_ServerInfo demonstrates how to use Admin.Instance.ServerInfo.
func TestExampleService_GetTableStats() {
	client, _ := misskey.NewClientWithOptions(misskey.WithSimpleConfig("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN")))

	response, err := client.Admin().Instance().GetTableStats()
	if err != nil {
		log.Printf("[Admin/Instance/GetTableStats] %s", err)

		return
	}

	log.Printf("Table Status: %v", response)
}
