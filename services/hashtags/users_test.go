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
	"github.com/yitsushi/go-misskey/models"
	"github.com/yitsushi/go-misskey/services/hashtags"
	"github.com/yitsushi/go-misskey/test"
)

func TestService_Users(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/hashtags/users",
		RequestData:  &hashtags.UsersRequest{},
		ResponseFile: "users.json",
		StatusCode:   http.StatusOK,
	})

	users, err := client.Hashtags().Users(hashtags.UsersRequest{
		Tag:   "vim",
		Sort:  hashtags.SortUsersByFollowers.Descending(),
		Limit: 7,
	})
	if !assert.NoError(t, err) {
		return
	}

	assert.Len(t, users, 7)
}

func TestService_Users_auth(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/hashtags/users",
		RequestData:  &hashtags.UsersRequest{},
		ResponseFile: "users_auth.json",
		StatusCode:   http.StatusOK,
	})

	users, err := client.Hashtags().Users(hashtags.UsersRequest{
		Tag:   "vim",
		Sort:  hashtags.SortUsersByFollowers.Descending(),
		Limit: 7,
	})
	if !assert.NoError(t, err) {
		return
	}

	assert.Len(t, users, 7)
}

func TestUsersRequest_Validate(t *testing.T) {
	test.ValidateRequests(
		t,
		[]core.BaseRequest{
			hashtags.UsersRequest{},
			hashtags.UsersRequest{
				Sort: hashtags.SortUsersByFollowers.Descending(),
			},
			hashtags.UsersRequest{
				Sort: hashtags.SortUsersByFollowers.Descending(),
				Tag:  "asd",
			},
		},
		[]core.BaseRequest{},
	)
}

func ExampleService_Users() {
	client, _ := misskey.NewClientWithOptions(misskey.WithSimpleConfig("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN")))
	client.LogLevel(logrus.DebugLevel)

	users, err := client.Hashtags().Users(hashtags.UsersRequest{
		Tag:    "vim",
		Limit:  20,
		State:  hashtags.DefaultState,
		Origin: models.OriginCombined,
		Sort:   hashtags.SortUsersByFollowers.Descending(),
	})
	if err != nil {
		log.Printf("[Hashtags] Error happened: %s", err)

		return
	}

	for _, user := range users {
		log.Println(user.Name)
	}
}
