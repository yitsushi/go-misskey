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

func TestService_UpdateRemoteUser(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/admin/update-remote-user",
		RequestData:  &federation.UpdateRemoteUserRequest{},
		ResponseFile: "empty",
		StatusCode:   http.StatusNoContent,
	})

	err := client.Admin().Federation().UpdateRemoteUser(federation.UpdateRemoteUserRequest{
		UserID: "1111",
	})

	assert.NoError(t, err)
}

func TestUpdateRemoteUserRequest_Validate(t *testing.T) {
	test.ValidateRequests(
		t,
		[]core.BaseRequest{
			federation.UpdateRemoteUserRequest{},
		},
		[]core.BaseRequest{
			federation.UpdateRemoteUserRequest{
				UserID: "1111",
			},
		},
	)
}

func ExampleService_UpdateRemoteUser() {
	client := misskey.NewClient("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN"))

	err := client.Admin().Federation().UpdateRemoteUser(federation.UpdateRemoteUserRequest{
		UserID: "1111",
	})
	if err != nil {
		log.Printf("[Admin/Federation] %s", err)
	}
}
