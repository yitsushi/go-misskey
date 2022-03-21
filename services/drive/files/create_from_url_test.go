package files_test

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/services/drive/files"
	"github.com/yitsushi/go-misskey/test"
)

type MockCreateFromURLHTTPClient struct{}

func (c *MockCreateFromURLHTTPClient) Do(req *http.Request) (*http.Response, error) {
	return &http.Response{
		Status:     "200 OK",
		StatusCode: 200,
		Proto:      "HTTP/1.0",
		ProtoMajor: 1,
		ProtoMinor: 0,
		Header:     http.Header{},
		Body:       io.NopCloser(strings.NewReader("example content, ignore me")),
	}, nil
}

func TestService_CreateFromURL(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/drive/files/create",
		RequestData:  &files.CreateRequest{},
		ResponseFile: "file.json",
		StatusCode:   http.StatusOK,
		Type:         test.IgnoreMockType,
	})

	response, err := client.Drive().File().CreateFromURL(files.CreateFromURLOptions{
		Name:           "file-name",
		FolderID:       "fancy-folder-id",
		URL:            "https://url.todownload.tld/file.jpg",
		DownloadClient: &MockCreateFromURLHTTPClient{},
	})
	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, "file-name", core.StringValue(response.Name))
}

func ExampleService_CreateFromURL() {
	client, _ := misskey.NewClientWithOptions(misskey.WithSimpleConfig("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN")))

	file, err := client.Drive().File().CreateFromURL(files.CreateFromURLOptions{
		Name:     "test-filename",
		FolderID: "8dmwisynnu",
		URL:      "https://www.wallpaperup.com/uploads/wallpapers/2014/01/23/235641/862478b1ad52546192af60ff03efbde9-700.jpg", //nolint:lll
	})
	if err != nil {
		log.Printf("[Drive/File/CreateFromURL] %s", err)

		return
	}

	log.Printf("[Drive/File/CreateFromURL] %s uploaded.", core.StringValue(file.Name))
}
