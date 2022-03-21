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

func TestService_Transfer(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/users/groups/transfer",
		RequestData:  &groups.TransferRequest{},
		ResponseFile: "show.json",
		StatusCode:   http.StatusOK,
	})

	group, err := client.Users().Groups().Transfer("93tyd132e7", "83sv4lyx22")
	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, "Test", group.Name)
}

func TestTransferRequest_Validate(t *testing.T) {
	test.ValidateRequests(
		t,
		[]core.BaseRequest{
			groups.TransferRequest{},
			groups.TransferRequest{GroupID: "93tyd132e7"},
			groups.TransferRequest{UserID: "93tyd132e7"},
		},
		[]core.BaseRequest{
			groups.TransferRequest{GroupID: "93tyd132e7", UserID: "83sv4lyx22"},
		},
	)
}

func ExampleService_Transfer() {
	client := misskey.NewClient("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN"))

	group, err := client.Users().Groups().Transfer("93tyd132e7", "83sv4lyx22")
	if err != nil {
		log.Printf("[Users/Groups/Transfer] %s", err)

		return
	}

	log.Printf("[Users/Groups/Transfer] %s", group.Name)
}
