package moderation_test

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/services/admin/moderation"
	"github.com/yitsushi/go-misskey/test"
)

func TestService_UserReports(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/admin/abuse-user-reports",
		RequestData:  &moderation.UserReportsRequest{},
		ResponseFile: "user_reports.json",
		StatusCode:   http.StatusOK,
	})

	response, err := client.Admin().Moderation().UserReports(moderation.UserReportsRequest{})

	if !assert.NoError(t, err) {
		return
	}

	assert.Len(t, response, 2)
}

func TestUserReportsRequest_Validate(t *testing.T) {
	test.ValidateRequests(
		t,
		[]core.BaseRequest{},
		[]core.BaseRequest{
			moderation.UserReportsRequest{},
		},
	)
}

func ExampleService_UserReports() {
	client := misskey.NewClient("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN"))

	response, err := client.Admin().Moderation().UserReports(moderation.UserReportsRequest{})
	if err != nil {
		log.Printf("[Admin/Moderation] %s", err)

		return
	}

	for _, item := range response {
		log.Printf("[Admin/Moderation] %s", item.Comment)
	}
}
