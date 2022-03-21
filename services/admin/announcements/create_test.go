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

func TestService_Create(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/admin/announcements/create",
		RequestData:  &announcements.CreateRequest{},
		ResponseFile: "create.json",
		StatusCode:   http.StatusOK,
	})

	response, err := client.Admin().Announcements().Create(announcements.CreateRequest{
		Title: "title",
		Text:  "text",
	})
	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, "8evhh5z4ni", core.StringValue(response.ID))
}

func TestCreateRequest_Validate(t *testing.T) {
	test.ValidateRequests(
		t,
		[]core.BaseRequest{
			announcements.CreateRequest{},
			announcements.CreateRequest{Title: "asd"},
			announcements.CreateRequest{
				Title:    "asd",
				Text:     "text",
				ImageURL: core.NewString(""),
			},
		},
		[]core.BaseRequest{
			announcements.CreateRequest{
				Title: "asd",
				Text:  "text",
			},
			announcements.CreateRequest{
				Title:    "asd",
				Text:     "text",
				ImageURL: core.NewString("image"),
			},
		},
	)
}

func ExampleService_Create() {
	client, _ := misskey.NewClientWithOptions(misskey.WithSimpleConfig("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN")))

	response, err := client.Admin().Announcements().Create(announcements.CreateRequest{
		Title: "New Announcement",
		Text:  "Because we can do it!",
	})
	if err != nil {
		log.Printf("[Admin/Announcements] %s", err)

		return
	}

	log.Printf("[Admin/Announcements] %s", *response.ID)
}
