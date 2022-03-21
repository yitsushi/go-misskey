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

func TestService_Files(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/admin/drive/files",
		RequestData:  &drive.FilesRequest{},
		ResponseFile: "files.json",
		StatusCode:   http.StatusOK,
	})

	response, err := client.Admin().Drive().Files(drive.FilesRequest{})

	if !assert.NoError(t, err) {
		return
	}

	assert.Len(t, response, 10)
}

func TestFilesRequest_Validate(t *testing.T) {
	test.ValidateRequests(
		t,
		[]core.BaseRequest{},
		[]core.BaseRequest{
			drive.FilesRequest{},
		},
	)
}

func ExampleService_Files() {
	client, _ := misskey.NewClientWithOptions(misskey.WithSimpleConfig("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN")))

	response, err := client.Admin().Drive().Files(drive.FilesRequest{})
	if err != nil {
		log.Printf("[Admin/Drive/Files] %s", err)

		return
	}

	for _, item := range response {
		log.Printf("[Admin/Drive/Files] %s", item.URL)
	}
}
