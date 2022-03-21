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

func TestService_Clean(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/admin/drive/cleanup",
		RequestData:  &drive.CleanRequest{},
		ResponseFile: "empty",
		StatusCode:   http.StatusNoContent,
	})

	err := client.Admin().Drive().Clean()

	if !assert.NoError(t, err) {
		return
	}
}

func TestCleanRequest_Validate(t *testing.T) {
	test.ValidateRequests(
		t,
		[]core.BaseRequest{},
		[]core.BaseRequest{
			drive.CleanRequest{},
		},
	)
}

func ExampleService_Clean() {
	client, _ := misskey.NewClientWithOptions(misskey.WithSimpleConfig("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN")))

	err := client.Admin().Drive().Clean()
	if err != nil {
		log.Printf("[Admin/Drive/Clean] %s", err)

		return
	}
}
