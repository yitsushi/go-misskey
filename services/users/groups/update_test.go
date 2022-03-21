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

func TestService_Update(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/users/groups/update",
		RequestData:  &groups.UpdateRequest{},
		ResponseFile: "show.json",
		StatusCode:   http.StatusOK,
	})

	group, err := client.Users().Groups().Update("93tyd132e7", "New Name")
	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, "Test", group.Name)
}

func TestUpdateRequest_Validate(t *testing.T) {
	test.ValidateRequests(
		t,
		[]core.BaseRequest{
			groups.UpdateRequest{},
			groups.UpdateRequest{GroupID: "93tyd132e7"},
			groups.UpdateRequest{Name: "new name"},
		},
		[]core.BaseRequest{
			groups.UpdateRequest{GroupID: "93tyd132e7", Name: "new name"},
		},
	)
}

func ExampleService_Update() {
	client, _ := misskey.NewClientWithOptions(misskey.WithSimpleConfig("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN")))

	group, err := client.Users().Groups().Update("93tyd132e7", "New Name")
	if err != nil {
		log.Printf("[Users/Groups/Update] %s", err)

		return
	}

	log.Printf("[Users/Groups/Update] %s", group.Name)
}
