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

func TestService_Instances(t *testing.T) {
	mockClient := test.SimpleMockEndpoint(&test.SimpleMockOptions{
		Endpoint:     "/api/federation/instances",
		RequestData:  &federation.InstancesRequest{},
		ResponseFile: "instances.json",
		StatusCode:   http.StatusOK,
	})

	client := misskey.NewClient("https://localhost", "thisistoken")
	client.HTTPClient = mockClient

	instances, err := client.Federation().Instances(federation.InstancesRequest{
		Limit: 10,
		Host:  "slippy.xyz",
	})
	if !assert.NoError(t, err) {
		return
	}

	assert.Len(t, instances, 2)
}

func TestInstanceRequest_Validate(t *testing.T) {
	testCase := federation.InstancesRequest{}
	assert.Error(t, testCase.Validate())

	testCase = federation.InstancesRequest{
		Host: "slippy.xyz",
	}
	assert.Error(t, testCase.Validate())

	testCase = federation.InstancesRequest{
		Host:  "slippy.xyz",
		Limit: 10,
	}
	assert.NoError(t, testCase.Validate())
}

func ExampleService_Instances() {
	client := misskey.NewClient("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN"))

	resp, err := client.Federation().Instances(federation.InstancesRequest{
		Limit: 100,
		Host:  "slippy.xyz",
	})
	if err != nil {
		log.Printf("[Federation/Instances] %s", err)

		return
	}

	log.Printf("[Federation/Instances] %v listed", resp)
}
