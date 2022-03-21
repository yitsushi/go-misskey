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

func TestService_Moderation(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/admin/show-moderation-logs",
		RequestData:  &logs.ModerationRequest{},
		ResponseFile: "moderation.json",
		StatusCode:   http.StatusOK,
	})

	response, err := client.Admin().Logs().Moderation()

	if !assert.NoError(t, err) {
		return
	}

	assert.Len(t, response, 4)
}

func TestModerationRequest_Validate(t *testing.T) {
	test.ValidateRequests(
		t,
		[]core.BaseRequest{},
		[]core.BaseRequest{
			logs.ModerationRequest{},
			logs.ModerationRequest{Limit: 10},
			logs.ModerationRequest{SinceID: "1111"},
			logs.ModerationRequest{Limit: 10, SinceID: "1111"},
			logs.ModerationRequest{UntilID: "9999"},
			logs.ModerationRequest{Limit: 10, SinceID: "1111"},
			logs.ModerationRequest{Limit: 10, UntilID: "9999", SinceID: "1111"},
		},
	)
}

func ExampleService_Moderation() {
	client, _ := misskey.NewClientWithOptions(misskey.WithSimpleConfig("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN")))

	response, err := client.Admin().Logs().Moderation()
	if err != nil {
		log.Printf("[Admin/Logs] %s", err)

		return
	}

	for _, item := range response {
		log.Printf("<%s> %s = %v", item.UserID, item.Type, item.Info)
	}
}
