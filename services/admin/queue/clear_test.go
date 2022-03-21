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

func TestService_Clear(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/admin/queue/clear",
		RequestData:  &queue.ClearRequest{},
		ResponseFile: "empty",
		StatusCode:   http.StatusNoContent,
	})

	err := client.Admin().Queue().Clear()

	assert.NoError(t, err)
}

func TestClearRequest_Validate(t *testing.T) {
	test.ValidateRequests(
		t,
		[]core.BaseRequest{},
		[]core.BaseRequest{
			queue.ClearRequest{},
		},
	)
}

func ExampleService_Clear() {
	client, _ := misskey.NewClientWithOptions(misskey.WithSimpleConfig("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN")))

	err := client.Admin().Queue().Clear()
	if err != nil {
		log.Printf("[Admin/Queue] %s", err)
	}
}
