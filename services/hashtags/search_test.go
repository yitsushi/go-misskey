package hashtags_test

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/services/hashtags"
	"github.com/yitsushi/go-misskey/test"
)

func TestService_Search(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/hashtags/search",
		RequestData:  &hashtags.SearchRequest{},
		ResponseFile: "search.json",
		StatusCode:   http.StatusOK,
	})

	tags, err := client.Hashtags().Search(hashtags.SearchRequest{
		Query: "hack%",
	})
	if !assert.NoError(t, err) {
		return
	}

	assert.Len(t, tags, 10)
}

func TestSearchRequest_Validate(t *testing.T) {
	test.ValidateRequests(
		t,
		[]core.BaseRequest{
			hashtags.SearchRequest{},
		},
		[]core.BaseRequest{},
	)
}

func ExampleService_Search() {
	client, _ := misskey.NewClientWithOptions(misskey.WithSimpleConfig("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN")))
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
