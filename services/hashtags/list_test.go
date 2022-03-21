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

func TestService_List(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/hashtags/list",
		RequestData:  &hashtags.ListRequest{},
		ResponseFile: "list.json",
		StatusCode:   http.StatusOK,
	})

	tags, err := client.Hashtags().List(hashtags.ListRequest{
		Sort: hashtags.SortTagsByAttachedUsers.Ascending(),
	})
	if !assert.NoError(t, err) {
		return
	}

	assert.Len(t, tags, 10)
}

func TestListRequest_Validate(t *testing.T) {
	test.ValidateRequests(
		t,
		[]core.BaseRequest{
			hashtags.ListRequest{},
		},
		[]core.BaseRequest{},
	)
}

func ExampleService_List() {
	client, _ := misskey.NewClientWithOptions(misskey.WithSimpleConfig("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN")))
	client.LogLevel(logrus.DebugLevel)

	tags, err := client.Hashtags().List(hashtags.ListRequest{
		Limit: 10,
		Sort:  hashtags.SortTagsByAttachedUsers.Ascending(),
	})
	if err != nil {
		log.Printf("[Hashtags] Error happened: %s", err)

		return
	}

	for _, tag := range tags {
		log.Printf(" - %s", tag.Tag)
	}
}
