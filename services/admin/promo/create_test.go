package promo_test

import (
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/services/admin/promo"
	"github.com/yitsushi/go-misskey/test"
)

func TestService_Create(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/admin/promo/create",
		RequestData:  &promo.CreateRequest{},
		ResponseFile: "create.json",
		StatusCode:   http.StatusOK,
	})

	err := client.Admin().Promo().Create(promo.CreateRequest{
		NoteID:    "noteID",
		ExpiresAt: time.Now().Add(86400 * 24 * time.Hour).Unix(),
	})
	if !assert.NoError(t, err) {
		return
	}
}

func TestService_Create_Error(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/admin/promo/create",
		RequestData:  &promo.CreateRequest{},
		ResponseFile: "create.json",
		StatusCode:   http.StatusOK,
	})

	// success
	err := client.Admin().Promo().Create(promo.CreateRequest{
		NoteID:    "noteID",
		ExpiresAt: time.Now().Add(86400 * 24 * time.Hour).Unix(),
	})
	if !assert.NoError(t, err) {
		return
	}
}

func TestPromoRequest_Validate(t *testing.T) {
	test.ValidateRequests(
		t,
		[]core.BaseRequest{
			promo.CreateRequest{},
		},
		[]core.BaseRequest{
			promo.CreateRequest{NoteID: "8zwxx3cpy7", ExpiresAt: 0},
			promo.CreateRequest{NoteID: "8zwxx3cpy8", ExpiresAt: 0},
			promo.CreateRequest{NoteID: "8zwxx3cpy9", ExpiresAt: time.Now().Add(86400 * 24 * time.Hour).Unix()},
		},
	)
}

func ExampleService_Create() {
	client, _ := misskey.NewClientWithOptions(misskey.WithSimpleConfig("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN")))

	err := client.Admin().Promo().Create(promo.CreateRequest{
		NoteID:    "8dsk7x47y3",
		ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
	})
	if err != nil {
		log.Printf("[Admin/Promo/Create] %s", err)

		return
	}
}
