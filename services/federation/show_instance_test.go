package federation_test

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/services/federation"
	"github.com/yitsushi/go-misskey/test"
)

func TestService_ShowInstance(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/federation/show-instance",
		RequestData:  &federation.ShowInstanceRequest{},
		ResponseFile: "show_instance.json",
		StatusCode:   http.StatusOK,
	})

	instance, err := client.Federation().ShowInstance(federation.ShowInstanceRequest{
		Host: "slippy.xyz",
	})
	if !assert.NoError(t, err) {
		return
	}

	assert.NotEmpty(t, instance)
	// Do a random check to ensure marshalling went okay.
	assert.Equal(t, "12.50.0", instance.SoftwareVersion)
}

func TestShowInstanceRequest_Validate(t *testing.T) {
	test.ValidateRequests(
		t,
		[]core.BaseRequest{
			federation.ShowInstanceRequest{},
		},
		[]core.BaseRequest{},
	)
}

func ExampleService_ShowInstance() {
	client := misskey.NewClient("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN"))

	resp, err := client.Federation().ShowInstance(federation.ShowInstanceRequest{
		Host: "slippy.xyz",
	})
	if err != nil {
		log.Printf("[Federation/ShowInstance] %s", err)

		return
	}

	log.Printf("[Federation/ShowInstance] %v listed", resp)
}
