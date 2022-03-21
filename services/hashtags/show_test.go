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

func TestService_Show(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/hashtags/show",
		RequestData:  &hashtags.ShowRequest{},
		ResponseFile: "show.json",
		StatusCode:   http.StatusOK,
	})

	tag, err := client.Hashtags().Show("hacktoberfest")
	if !assert.NoError(t, err) {
		return
	}

	assert.EqualValues(t, 10, tag.MentionedUsersCount)
	assert.EqualValues(t, 7, tag.MentionedLocalUsersCount)
	assert.EqualValues(t, 3, tag.MentionedRemoteUsersCount)

	assert.EqualValues(t, 3, tag.AttachedUsersCount)
	assert.EqualValues(t, 1, tag.AttachedLocalUsersCount)
	assert.EqualValues(t, 2, tag.AttachedRemoteUsersCount)
}

func ExampleService_Show() {
	client, _ := misskey.NewClientWithOptions(misskey.WithSimpleConfig("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN")))
	client.LogLevel(logrus.DebugLevel)

	tag, err := client.Hashtags().Show("hacktoberfest")
	if err != nil {
		log.Printf("[Hashtags] Error happened: %s", err)

		return
	}

	log.Println(tag)
}
