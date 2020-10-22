package antennas_test

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/services/antennas"
	"github.com/yitsushi/go-misskey/test"
)

func TestService_Delete(t *testing.T) {
	mockClient := test.SimpleMockEndpoint(&test.SimpleMockOptions{
		Endpoint:     "/api/antennas/delete",
		RequestData:  &antennas.DeleteRequest{},
		ResponseFile: "empty",
		StatusCode:   http.StatusNoContent,
	})

	client := misskey.NewClient("https://localhost", "thisistoken")
	client.HTTPClient = mockClient

	err := client.Antennas().Delete("test")
	if !assert.NoError(t, err) {
		return
	}
}

func ExampleService_Delete() {
	client := misskey.NewClient("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN"))

	err := client.Antennas().Delete("antenna-id")
	if err != nil {
		log.Printf("[Antennas/Delete] %s", err)

		return
	}

	log.Println("[Antennas/Delete] Done without errors")
}
