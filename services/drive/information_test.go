package drive_test

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/services/drive"
	"github.com/yitsushi/go-misskey/test"
)

func TestService_Information(t *testing.T) {
	mockClient := test.SimpleMockEndpoint(&test.SimpleMockOptions{
		Endpoint:     "/api/drive",
		RequestData:  &drive.InformationRequest{},
		ResponseFile: "information.json",
		StatusCode:   http.StatusOK,
	})

	client := misskey.NewClient("https://localhost", "thisistoken")
	client.HTTPClient = mockClient

	info, err := client.Drive().Information()
	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, uint64(2147483648), info.Capacity.Bytes())
	assert.Equal(t, uint64(343367058), info.Usage.Bytes())
}

func ExampleService_Information() {
	client := misskey.NewClient("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN"))

	info, err := client.Drive().Information()
	if err != nil {
		log.Printf("[Drive/Information] %s", err)

		return
	}

	log.Printf("[Drive/Information] Capacity: %.2f", info.Capacity.Gigabytes())
}
