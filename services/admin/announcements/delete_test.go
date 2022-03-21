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

func TestService_Delete(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/admin/announcements/delete",
		RequestData:  &announcements.DeleteRequest{},
		ResponseFile: "empty",
		StatusCode:   http.StatusNoContent,
	})

	err := client.Admin().Announcements().Delete("8d44utwtj6")

	assert.NoError(t, err)
}

func TestDeleteRequest_Validate(t *testing.T) {
	test.ValidateRequests(
		t,
		[]core.BaseRequest{
			announcements.DeleteRequest{},
			announcements.DeleteRequest{ID: ""},
		},
		[]core.BaseRequest{
			announcements.DeleteRequest{ID: "8d44utwtj6"},
		},
	)
}

func ExampleService_Delete() {
	client, _ := misskey.NewClientWithOptions(misskey.WithSimpleConfig("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN")))

	err := client.Admin().Announcements().Delete("8d44utwtj6")
	if err != nil {
		log.Printf("[Admin/Announcements] %s", err)

		return
	}
}
