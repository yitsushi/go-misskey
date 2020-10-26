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

func TestService_Trend(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/hashtags/trend",
		RequestData:  &hashtags.TrendRequest{},
		ResponseFile: "trend.json",
		StatusCode:   http.StatusOK,
	})

	trend, err := client.Hashtags().Trend()
	if !assert.NoError(t, err) {
		return
	}

	assert.Len(t, trend, 5)

	one := trend[0]

	assert.Equal(t, "caturday", one.Tag)
	assert.EqualValues(t, 2, one.UsersCount)
	assert.Len(t, one.Chart, 20)
}

func ExampleService_Trend() {
	client := misskey.NewClient("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN"))
	client.LogLevel(logrus.DebugLevel)

	trend, err := client.Hashtags().Trend()
	if err != nil {
		log.Printf("[Hashtags] Error happened: %s", err)

		return
	}

	for _, tag := range trend {
		log.Println(tag.Tag)
	}
}
