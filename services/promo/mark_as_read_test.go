package promo_test

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/services/promo"
	"github.com/yitsushi/go-misskey/test"
)

func TestService_MarkAsRead(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/promo/read",
		RequestData:  &promo.MarkAsReadRequest{},
		ResponseFile: "empty",
		StatusCode:   http.StatusNoContent,
	})

	err := client.Promo().MarkAsRead("8dsk7x47y3")
	if !assert.NoError(t, err) {
		return
	}
}

func ExampleService_MarkAsRead() {
	client, _ := misskey.NewClientWithOptions(misskey.WithSimpleConfig("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN")))

	err := client.Promo().MarkAsRead("8dsk7x47y3")
	if err != nil {
		log.Printf("[Promo/MarkAsRead] %s", err)

		return
	}
}
