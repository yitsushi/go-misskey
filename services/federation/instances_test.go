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

func TestService_Instances(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/federation/instances",
		RequestData:  &federation.InstancesRequest{},
		ResponseFile: "instances.json",
		StatusCode:   http.StatusOK,
	})

	instances, err := client.Federation().Instances(federation.InstancesRequest{
		Limit: 2,
		Host:  "slippy.xyz",
	})
	if !assert.NoError(t, err) {
		return
	}

	assert.Len(t, instances, 2)
}

func TestInstanceRequest_Validate(t *testing.T) {
	test.ValidateRequests(
		t,
		[]core.BaseRequest{
			federation.InstancesRequest{},
			federation.InstancesRequest{Host: "slippy.xyz"},
		},
		[]core.BaseRequest{},
	)
}

func ExampleService_Instances() {
	client := misskey.NewClient("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN"))

	resp, err := client.Federation().Instances(federation.InstancesRequest{
		Limit: 90,
		Host:  "slippy.xyz",
	})
	if err != nil {
		log.Printf("[Federation/Instances] %s", err)

		return
	}

	log.Printf("[Federation/Instances] %v listed", resp)
}
