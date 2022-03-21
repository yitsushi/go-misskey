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

func TestService_Add(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/admin/moderators/add",
		RequestData:  &moderators.AddRequest{},
		ResponseFile: "empty",
		StatusCode:   http.StatusNoContent,
	})

	err := client.Admin().Moderators().Add("88v9vu5nbu")

	assert.NoError(t, err)
}

func TestAddRequest_Validate(t *testing.T) {
	test.ValidateRequests(
		t,
		[]core.BaseRequest{
			moderators.AddRequest{},
		},
		[]core.BaseRequest{
			moderators.AddRequest{UserID: "asd"},
		},
	)
}

func ExampleService_Add() {
	client := misskey.NewClient("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN"))

	err := client.Admin().Moderators().Add("88v9vu5nbu")
	if err != nil {
		log.Printf("[Admin/Moderators/Add] %s", err)

		return
	}
}
