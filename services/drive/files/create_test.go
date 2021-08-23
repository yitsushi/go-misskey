package files_test

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/services/drive/files"
	"github.com/yitsushi/go-misskey/test"
)

func TestService_Create(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/drive/files/create",
		RequestData:  &files.CreateRequest{},
		ResponseFile: "file.json",
		StatusCode:   http.StatusOK,
		Type:         test.IgnoreMockType,
	})

	response, err := client.Drive().File().Create(files.CreateRequest{
		Name:        "file-name",
		FolderID:    "fancy-folder-id",
		IsSensitive: false,
		Force:       false,
		Content:     []byte("my content"),
	})
	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, "file-name", core.StringValue(response.Name))
}

func ExampleService_Create() {
	client := misskey.NewClient("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN"))

	fileContent := []byte{}

	file, err := client.Drive().File().Create(files.CreateRequest{
		FolderID:    "",
		Name:        "this is the name",
		IsSensitive: false,
		Force:       false,
		Content:     fileContent,
	})
	if err != nil {
		log.Printf("[Drive/File/Create] %s", err)

		return
	}

	log.Printf(
		"[Drive/File/Create] %s file uploaded. (%s)",
		core.StringValue(file.Name),
		file.ID,
	)
}
