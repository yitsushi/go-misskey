package instance_test

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/services/admin/instance"
	"github.com/yitsushi/go-misskey/test"
)

func TestGetIndexStats(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/admin/get-index-stats",
		RequestData:  &instance.GetIndexStatsRequest{},
		ResponseFile: "get-index-stats.json",
		StatusCode:   http.StatusOK,
	})

	response, err := client.Admin().Instance().GetIndexStats()

	require.NoError(t, err)
	require.Len(t, response, 518)
	assert.Equal(t, "CREATE UNIQUE INDEX pg_proc_proname_args_nsp_index ON pg_catalog.pg_proc USING btree (proname, proargtypes, pronamespace)", response[0].Indexdef)
	assert.Equal(t, "pg_proc_proname_args_nsp_index", response[0].Indexname)
	assert.Equal(t, "pg_catalog", response[0].Schemaname)
	assert.Equal(t, "pg_proc", response[0].Tablename)
	assert.Nil(t, response[0].Tablespace)
}

// ExampleService_GetIndexStats demonstrates how to use Admin.Instance.GetIndexStats.
func ExampleService_GetIndexStats() {
	client, _ := misskey.NewClientWithOptions(misskey.WithSimpleConfig("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN")))

	response, err := client.Admin().Instance().GetIndexStats()
	if err != nil {
		log.Printf("[Admin/Instance/GetIndexStats] %s", err)

		return
	}

	log.Printf("Table Status: %v", response)
}
