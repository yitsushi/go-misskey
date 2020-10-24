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

func TestService_Search(t *testing.T) {
	mockClient := test.SimpleMockEndpoint(&test.SimpleMockOptions{
		Endpoint:     "/api/hashtags/search",
		RequestData:  &hashtags.SearchRequest{},
		ResponseFile: "search.json",
		StatusCode:   http.StatusOK,
	})

	client := misskey.NewClient("https://localhost", "thisistoken")
	client.HTTPClient = mockClient

	tags, err := client.Hashtags().Search(hashtags.SearchRequest{
		Query: "hack%",
	})
	if !assert.NoError(t, err) {
		return
	}

	assert.Len(t, tags, 10)
}

func TestService_Search_missingQuery(t *testing.T) {
	mockClient := test.SimpleMockEndpoint(&test.SimpleMockOptions{
		Endpoint:     "/api/hashtags/search",
		RequestData:  &hashtags.SearchRequest{},
		ResponseFile: "search.json",
		StatusCode:   http.StatusOK,
	})

	client := misskey.NewClient("https://localhost", "thisistoken")
	client.HTTPClient = mockClient

	_, err := client.Hashtags().Search(hashtags.SearchRequest{
		Limit: 10,
	})

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "[Query] Undefined required field")
}

func ExampleService_Search() {
	client := misskey.NewClient("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN"))
	client.LogLevel(logrus.DebugLevel)

	tags, err := client.Hashtags().Search(hashtags.SearchRequest{
		Limit: 10,
		Query: "hack%",
	})
	if err != nil {
		log.Printf("[Hashtags] Error happened: %s", err)

		return
	}

	for _, tag := range tags {
		log.Printf(" - %s", tag)
	}
}
