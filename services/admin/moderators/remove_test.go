package moderators_test

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/services/admin/moderators"
	"github.com/yitsushi/go-misskey/test"
)

func TestService_Remove(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/admin/moderators/remove",
		RequestData:  &moderators.RemoveRequest{},
		ResponseFile: "empty",
		StatusCode:   http.StatusNoContent,
	})

	err := client.Admin().Moderators().Remove("88v9vu5nbu")

	assert.NoError(t, err)
}

func TestRemoveRequest_Validate(t *testing.T) {
	test.ValidateRequests(
		t,
		[]core.BaseRequest{
			moderators.RemoveRequest{},
		},
		[]core.BaseRequest{
			moderators.RemoveRequest{UserID: "asd"},
		},
	)
}

func ExampleService_Remove() {
	client := misskey.NewClient("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN"))

	err := client.Admin().Moderators().Remove("88v9vu5nbu")
	if err != nil {
		log.Printf("[Admin/Moderators/Remove] %s", err)

		return
	}
}
