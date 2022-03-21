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

func TestService_ResolveReport(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/admin/resolve-abuse-user-report",
		RequestData:  &moderation.ResolveReportRequest{},
		ResponseFile: "empty",
		StatusCode:   http.StatusNoContent,
	})

	err := client.Admin().Moderation().ResolveReport("8evj2lmh10")

	assert.NoError(t, err)
}

func TestResolveReportRequest_Validate(t *testing.T) {
	test.ValidateRequests(
		t,
		[]core.BaseRequest{
			moderation.ResolveReportRequest{},
		},
		[]core.BaseRequest{
			moderation.ResolveReportRequest{ID: "asd"},
		},
	)
}

func ExampleService_ResolveReport() {
	client, _ := misskey.NewClientWithOptions(misskey.WithSimpleConfig("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN")))

	err := client.Admin().Moderation().ResolveReport("8evj2lmh10")
	if err != nil {
		log.Printf("[Admin/Moderation] %s", err)

		return
	}
}
