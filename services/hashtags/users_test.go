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

func TestService_Users(t *testing.T) {
	mockClient := test.SimpleMockEndpoint(&test.SimpleMockOptions{
		Endpoint:     "/api/hashtags/users",
		RequestData:  &hashtags.UsersRequest{},
		ResponseFile: "users.json",
		StatusCode:   http.StatusOK,
	})

	client := misskey.NewClient("https://localhost", "thisistoken")
	client.HTTPClient = mockClient

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
	mockClient := test.SimpleMockEndpoint(&test.SimpleMockOptions{
		Endpoint:     "/api/hashtags/users",
		RequestData:  &hashtags.UsersRequest{},
		ResponseFile: "users_auth.json",
		StatusCode:   http.StatusOK,
	})

	client := misskey.NewClient("https://localhost", "thisistoken")
	client.HTTPClient = mockClient

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

func TestService_Users_missingSort(t *testing.T) {
	mockClient := test.SimpleMockEndpoint(&test.SimpleMockOptions{
		Endpoint:     "/api/hashtags/users",
		RequestData:  &hashtags.UsersRequest{},
		ResponseFile: "users.json",
		StatusCode:   http.StatusOK,
	})

	client := misskey.NewClient("https://localhost", "thisistoken")
	client.HTTPClient = mockClient

	_, err := client.Hashtags().Users(hashtags.UsersRequest{
		Limit: 7,
		Tag:   "vim",
	})

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "[Sort]")
}

func TestService_Users_missingTag(t *testing.T) {
	mockClient := test.SimpleMockEndpoint(&test.SimpleMockOptions{
		Endpoint:     "/api/hashtags/users",
		RequestData:  &hashtags.UsersRequest{},
		ResponseFile: "users.json",
		StatusCode:   http.StatusOK,
	})

	client := misskey.NewClient("https://localhost", "thisistoken")
	client.HTTPClient = mockClient

	_, err := client.Hashtags().Users(hashtags.UsersRequest{
		Limit: 7,
		Sort:  hashtags.SortUsersByFollowers.Descending(),
	})

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "[Tag]")
}

func TestService_Users_missingSortAndTag(t *testing.T) {
	mockClient := test.SimpleMockEndpoint(&test.SimpleMockOptions{
		Endpoint:     "/api/hashtags/users",
		RequestData:  &hashtags.UsersRequest{},
		ResponseFile: "users.json",
		StatusCode:   http.StatusOK,
	})

	client := misskey.NewClient("https://localhost", "thisistoken")
	client.HTTPClient = mockClient

	_, err := client.Hashtags().Users(hashtags.UsersRequest{
		Limit: 7,
	})

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "[Sort]")
}

func ExampleService_Users() {
	client := misskey.NewClient("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN"))
	client.LogLevel(logrus.DebugLevel)

	users, err := client.Hashtags().Users(hashtags.UsersRequest{
		Tag:    "vim",
		Limit:  20,
		State:  hashtags.DefaultState,
		Origin: hashtags.OriginCombined,
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
