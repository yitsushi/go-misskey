package logs_test

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/services/admin/logs"
	"github.com/yitsushi/go-misskey/test"
)

func TestService_Clear(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/admin/delete-logs",
		RequestData:  &logs.ClearRequest{},
		ResponseFile: "empty",
		StatusCode:   http.StatusNoContent,
	})

	err := client.Admin().Logs().Clear()

	assert.NoError(t, err)
}

func TestClearRequest_Validate(t *testing.T) {
	test.ValidateRequests(
		t,
		[]core.BaseRequest{},
		[]core.BaseRequest{
			logs.ClearRequest{},
		},
	)
}

func ExampleService_Clear() {
	client, _ := misskey.NewClientWithOptions(misskey.WithSimpleConfig("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN")))

	err := client.Admin().Logs().Clear()
	if err != nil {
		log.Printf("[Admin/Logs] %s", err)
	}
}
