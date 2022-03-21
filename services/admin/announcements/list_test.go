package announcements_test

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/services/admin/announcements"
	"github.com/yitsushi/go-misskey/test"
)

func TestService_List(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/admin/announcements/list",
		RequestData:  &announcements.ListRequest{},
		ResponseFile: "list.json",
		StatusCode:   http.StatusOK,
	})

	response, err := client.Admin().Announcements().List(announcements.ListRequest{
		Limit: 1,
	})
	if !assert.NoError(t, err) {
		return
	}

	assert.Len(t, response, 1)
}

func TestListRequest_Validate(t *testing.T) {
	test.ValidateRequests(
		t,
		[]core.BaseRequest{
			announcements.ListRequest{},
			announcements.ListRequest{Limit: 200},
		},
		[]core.BaseRequest{
			announcements.ListRequest{Limit: 10},
		},
	)
}

func ExampleService_List() {
	client, _ := misskey.NewClientWithOptions(misskey.WithSimpleConfig("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN")))

	response, err := client.Admin().Announcements().List(announcements.ListRequest{
		Limit: 10,
	})
	if err != nil {
		log.Printf("[Admin/Announcements] %s", err)

		return
	}

	for _, item := range response {
		log.Printf("[Admin/Announcements] %s", *item.Title)
	}
}
