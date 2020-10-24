package hashtags_test

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/services/hashtags"
	"github.com/yitsushi/go-misskey/test"
)

func TestService_List(t *testing.T) {
	mockClient := test.SimpleMockEndpoint(&test.SimpleMockOptions{
		Endpoint:     "/api/hashtags/list",
		RequestData:  &hashtags.ListRequest{},
		ResponseFile: "list.json",
		StatusCode:   http.StatusOK,
	})

	client := misskey.NewClient("https://localhost", "thisistoken")
	client.HTTPClient = mockClient

	tags, err := client.Hashtags().List(&hashtags.ListOptions{
		Sort: hashtags.SortAttachedLocalUsers.Ascending(),
	})
	if !assert.NoError(t, err) {
		return
	}

	assert.Len(t, tags, 10)
}

func TestService_List_missingSort(t *testing.T) {
	mockClient := test.SimpleMockEndpoint(&test.SimpleMockOptions{
		Endpoint:     "/api/hashtags/list",
		RequestData:  &hashtags.ListRequest{},
		ResponseFile: "list.json",
		StatusCode:   http.StatusOK,
	})

	client := misskey.NewClient("https://localhost", "thisistoken")
	client.HTTPClient = mockClient

	_, err := client.Hashtags().List(&hashtags.ListOptions{
		Limit: 10,
	})

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "endpoint: Sort")
}

func ExampleService_List() {
	client := misskey.NewClient("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN"))
	client.LogLevel(logrus.DebugLevel)

	tags, err := client.Hashtags().List(&hashtags.ListOptions{
		Limit: 10,
		Sort:  hashtags.SortAttachedLocalUsers.Ascending(),
	})
	if err != nil {
		log.Printf("[Hashtags] Error happened: %s", err)

		return
	}

	for _, tag := range tags {
		log.Printf(" - %s", tag.Tag)
	}
}
