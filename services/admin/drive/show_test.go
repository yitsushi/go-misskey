package drive_test

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/services/admin/drive"
	"github.com/yitsushi/go-misskey/test"
)

func TestService_Show(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/admin/drive/show-file",
		RequestData:  &drive.ShowRequest{},
		ResponseFile: "file.json",
		StatusCode:   http.StatusOK,
	})

	response, err := client.Admin().Drive().Show(drive.ShowRequest{})

	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, "8y1rwxxkk9", response.ID)
}

func TestShowRequest_Validate(t *testing.T) {
	test.ValidateRequests(
		t,
		[]core.BaseRequest{},
		[]core.BaseRequest{
			drive.ShowRequest{},
		},
	)
}

func ExampleService_Show() {
	client, _ := misskey.NewClientWithOptions(misskey.WithSimpleConfig("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN")))

	response, err := client.Admin().Drive().Show(drive.ShowRequest{})
	if err != nil {
		log.Printf("[Admin/Drive/Show] %s", err)

		return
	}

	log.Printf("[Admin/Drive/Show] %s", response.URL)
}
