package federation_test

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/services/admin/federation"
	"github.com/yitsushi/go-misskey/test"
)

func TestService_UpdateInstance(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/admin/federation/update-instance",
		RequestData:  &federation.UpdateInstanceRequest{},
		ResponseFile: "empty",
		StatusCode:   http.StatusNoContent,
	})

	err := client.Admin().Federation().UpdateInstance(federation.UpdateInstanceRequest{
		Host:        "misskey.io",
		IsSuspended: false,
	})

	assert.NoError(t, err)
}

func TestUpdateInstanceRequest_Validate(t *testing.T) {
	test.ValidateRequests(
		t,
		[]core.BaseRequest{
			federation.UpdateInstanceRequest{},
			federation.UpdateInstanceRequest{IsSuspended: false},
		},
		[]core.BaseRequest{
			federation.UpdateInstanceRequest{
				Host:        "misskey.io",
				IsSuspended: false,
			},
			federation.UpdateInstanceRequest{
				Host: "misskey.io",
			},
		},
	)
}

func ExampleService_UpdateInstance() {
	client, _ := misskey.NewClientWithOptions(misskey.WithSimpleConfig("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN")))

	err := client.Admin().Federation().UpdateInstance(federation.UpdateInstanceRequest{
		Host:        "misskey.io",
		IsSuspended: false,
	})
	if err != nil {
		log.Printf("[Admin/Federation] %s", err)
	}
}
