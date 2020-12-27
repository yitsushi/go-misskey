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

func TestService_Server(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/admin/logs",
		RequestData:  &logs.ServerRequest{},
		ResponseFile: "server.json",
		StatusCode:   http.StatusOK,
	})

	response, err := client.Admin().Logs().Server()

	if !assert.NoError(t, err) {
		return
	}

	assert.Len(t, response, 9)
}

func TestServerRequest_Validate(t *testing.T) {
	test.ValidateRequests(
		t,
		[]core.BaseRequest{},
		[]core.BaseRequest{
			logs.ServerRequest{},
			logs.ServerRequest{Limit: 10},
			logs.ServerRequest{Domain: "chart"},
			logs.ServerRequest{Limit: 10, Domain: "chart"},
			logs.ServerRequest{Level: "info"},
			logs.ServerRequest{Limit: 10, Level: "info"},
			logs.ServerRequest{Limit: 10, Domain: "chart", Level: "info"},
		},
	)
}

func ExampleService_Server() {
	client := misskey.NewClient("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN"))

	response, err := client.Admin().Logs().Server()
	if err != nil {
		log.Printf("[Admin/Logs] %s", err)

		return
	}

	for _, item := range response {
		log.Printf("<%s> %v = %s", item.Level, item.Domain, item.Message)
	}
}
