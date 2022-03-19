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

func TestService_CleanRemote(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/admin/drive/clean-remote-files",
		RequestData:  &drive.CleanRemoteRequest{},
		ResponseFile: "empty",
		StatusCode:   http.StatusNoContent,
	})

	err := client.Admin().Drive().CleanRemote()

	if !assert.NoError(t, err) {
		return
	}
}

func TestCleanRemoteRequest_Validate(t *testing.T) {
	test.ValidateRequests(
		t,
		[]core.BaseRequest{},
		[]core.BaseRequest{
			drive.CleanRemoteRequest{},
		},
	)
}

func ExampleService_CleanRemote() {
	client := misskey.NewClient("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN"))

	err := client.Admin().Drive().CleanRemote()
	if err != nil {
		log.Printf("[Admin/Drive/CleanRemote] %s", err)

		return
	}
}
