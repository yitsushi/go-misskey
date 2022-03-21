package groups_test

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/services/users/groups"
	"github.com/yitsushi/go-misskey/test"
)

func TestService_Delete(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/users/groups/delete",
		RequestData:  &groups.DeleteRequest{},
		ResponseFile: "empty",
		StatusCode:   http.StatusNoContent,
	})

	err := client.Users().Groups().Delete("8y4nlhyx3v")
	if !assert.NoError(t, err) {
		return
	}
}

func TestDeleteRequest_Validate(t *testing.T) {
	test.ValidateRequests(
		t,
		[]core.BaseRequest{
			groups.DeleteRequest{},
		},
		[]core.BaseRequest{
			groups.DeleteRequest{GroupID: "8y4nlhyx3v"},
		},
	)
}

func ExampleService_Delete() {
	client, _ := misskey.NewClientWithOptions(misskey.WithSimpleConfig("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN")))

	err := client.Users().Groups().Delete("8y4nlhyx3v")
	if err != nil {
		log.Printf("[Users/Groups/Delete] %s", err)

		return
	}
}
