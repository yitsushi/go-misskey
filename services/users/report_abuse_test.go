package users_test

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/services/users"
	"github.com/yitsushi/go-misskey/test"
)

func TestReportAbuseRequest_Validate(t *testing.T) {
	test.ValidateRequests(
		t,
		[]core.BaseRequest{
			users.ReportAbuseRequest{},
			users.ReportAbuseRequest{Comment: "A comment without a user-id."},
			users.ReportAbuseRequest{UserID: "8y1nj3wzmz"},
		},
		[]core.BaseRequest{
			users.ReportAbuseRequest{UserID: "8gf082lv8f", Comment: "A comment with a user-id."},
		},
	)
}

func TestService_ReportAbuse(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/users/report-abuse",
		RequestData:  &users.ReportAbuseRequest{},
		ResponseFile: "empty",
		StatusCode:   http.StatusNoContent,
	})

	err := client.Users().ReportAbuse("88v9vu5nbu", "A comment.")

	assert.NoError(t, err)
}

// ExampleService_ReportAbuse is an example of how to use the ReportAbuse method to file a report.
func ExampleService_ReportAbuse() {
	client, _ := misskey.NewClientWithOptions(misskey.WithSimpleConfig("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN")))

	err := client.Users().ReportAbuse("88v9vu5nbu", "A comment.")
	if err != nil {
		log.Printf("[Users/ReportAbuse] %s", err)

		return
	}
}
