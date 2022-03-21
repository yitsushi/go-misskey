package federation_test

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/services/admin/federation"
	"github.com/yitsushi/go-misskey/test"
)

func TestService_DeleteAllFiles(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/admin/federation/delete-all-files",
		RequestData:  &federation.DeleteAllFilesRequest{},
		ResponseFile: "empty",
		StatusCode:   http.StatusNoContent,
	})

	err := client.Admin().Federation().DeleteAllFiles(federation.DeleteAllFilesRequest{
		Host: "quey.org",
	})

	assert.NoError(t, err)
}

func TestDeleteAllFilesRequest_Validate(t *testing.T) {
	test.ValidateRequests(
		t,
		[]core.BaseRequest{
			federation.DeleteAllFilesRequest{},
		},
		[]core.BaseRequest{
			federation.DeleteAllFilesRequest{
				Host: "quey.org",
			},
		},
	)
}

func ExampleService_DeleteAllFiles() {
	client, _ := misskey.NewClientWithOptions(misskey.WithSimpleConfig("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN")))

	err := client.Admin().Federation().DeleteAllFiles(federation.DeleteAllFilesRequest{
		Host: "quey.org",
	})
	if err != nil {
		log.Printf("[Admin/Federation] %s", err)
	}
}
