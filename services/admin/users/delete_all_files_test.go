package users_test

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/services/admin/users"
	"github.com/yitsushi/go-misskey/test"
)

func TestService_DeleteAllFiles(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/admin/delete-all-files-of-a-user",
		RequestData:  &users.DeleteAllFilesRequest{},
		ResponseFile: "empty",
		StatusCode:   http.StatusNoContent,
	})

	err := client.Admin().Users().DeleteAllFiles("88v9vu5nbu")
	assert.NoError(t, err)
}

func TestDeleteAllFilesRequest_Validate(t *testing.T) {
	test.ValidateRequests(
		t,
		[]core.BaseRequest{
			users.DeleteAllFilesRequest{},
		},
		[]core.BaseRequest{
			users.DeleteAllFilesRequest{UserID: "asd"},
		},
	)
}

func ExampleService_DeleteAllFiles() {
	client := misskey.NewClient("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN"))

	err := client.Admin().Users().DeleteAllFiles("88v9vu5nbu")
	if err != nil {
		log.Printf("[Admin/Users] %s", err)

		return
	}
}
