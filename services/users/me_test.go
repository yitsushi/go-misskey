package users_test

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/go-misskey"
	_ "github.com/yitsushi/go-misskey/services/users"
	"github.com/yitsushi/go-misskey/test"
)

func TestService_Me(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/i",
		ResponseFile: "i.json",
		StatusCode:   http.StatusOK,
	})

	user, err := client.Users().Me()
	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, user.ID, "9dr5dkiiby")
}

func ExampleService_Me() {
	client, _ := misskey.NewClientWithOptions(misskey.WithSimpleConfig("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN")))

	user, err := client.Users().Me()
	if err != nil {
		log.Printf("[Users/Me] %s", err)

		return
	}

	log.Printf("[Users/Me] <%s> %s", user.Username, user.Name)
}
