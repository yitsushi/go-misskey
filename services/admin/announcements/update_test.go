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

func TestService_Update(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/admin/announcements/update",
		RequestData:  &announcements.UpdateRequest{},
		ResponseFile: "empty",
		StatusCode:   http.StatusNoContent,
	})

	err := client.Admin().Accouncements().Update(announcements.UpdateRequest{
		ID:    "8d44utwtj6",
		Title: "New Title",
		Text:  "New text",
	})

	assert.NoError(t, err)
}

func TestUpdateRequest_Validate(t *testing.T) {
	test.ValidateRequests(
		t,
		[]core.BaseRequest{
			announcements.UpdateRequest{},
			announcements.UpdateRequest{ID: "asd"},
			announcements.UpdateRequest{ID: "asd", Title: "asd"},
			announcements.UpdateRequest{
				ID:       "asd",
				Title:    "asd",
				Text:     "asd",
				ImageURL: core.NewString(""),
			},
		},
		[]core.BaseRequest{
			announcements.UpdateRequest{
				ID:    "8d44utwtj6",
				Title: "title",
				Text:  "text",
			},
			announcements.UpdateRequest{
				ID:       "8d44utwtj6",
				Title:    "title",
				Text:     "text",
				ImageURL: core.NewString("asd"),
			},
		},
	)
}

func ExampleService_Update() {
	client := misskey.NewClient("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN"))

	err := client.Admin().Accouncements().Update(announcements.UpdateRequest{
		ID:    "8d44utwtj6",
		Title: "New Title",
		Text:  "New text",
	})
	if err != nil {
		log.Printf("[Admin/Announcements] %s", err)

		return
	}
}
