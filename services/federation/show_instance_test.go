package federation_test

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/services/federation"
	"github.com/yitsushi/go-misskey/test"
)

func TestService_ShowInstance(t *testing.T) {
	mockClient := test.SimpleMockEndpoint(&test.SimpleMockOptions{
		Endpoint:     "/api/federation/show-instance",
		RequestData:  &federation.ShowInstanceRequest{},
		ResponseFile: "show_instance.json",
		StatusCode:   http.StatusOK,
	})

	client := misskey.NewClient("https://localhost", "thisistoken")
	client.HTTPClient = mockClient

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
	testCase := federation.ShowInstanceRequest{}
	assert.Error(t, testCase.Validate())

	testCase = federation.ShowInstanceRequest{}
	assert.Error(t, testCase.Validate())

	testCase = federation.ShowInstanceRequest{
		Host: "slippy.xyz",
	}
	assert.NoError(t, testCase.Validate())
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
