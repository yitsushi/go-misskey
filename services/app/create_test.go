package app_test

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/core/permissions"
	"github.com/yitsushi/go-misskey/services/app"
	"github.com/yitsushi/go-misskey/test"
)

func TestService_Create(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/app/create",
		RequestData:  &app.CreateRequest{},
		ResponseFile: "app.json",
		StatusCode:   http.StatusOK,
	})

	resp, err := client.App().Create(app.CreateRequest{
		Name:        "test app",
		Description: "my test app",
		Permission: []permissions.Permission{
			permissions.Write(permissions.Account),
			permissions.Read(permissions.Account),
		},
	})
	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, "test app", resp.Name)
}

func TestCreateRequest_Validate(t *testing.T) {
	test.ValidateRequests(
		t,
		[]core.BaseRequest{
			app.CreateRequest{},
			app.CreateRequest{Name: "asd"},
			app.CreateRequest{Description: "asd"},
			app.CreateRequest{Permission: []permissions.Permission{"asd"}},
			app.CreateRequest{Name: "asd", Description: "asd"},
			app.CreateRequest{
				Name:       "asd",
				Permission: []permissions.Permission{"asd"},
			},
			app.CreateRequest{
				Description: "asd",
				Permission:  []permissions.Permission{"asd"},
			},
			app.CreateRequest{
				Name:        "asd",
				Description: "asd",
				Permission:  []permissions.Permission{},
			},
		},
		[]core.BaseRequest{
			app.CreateRequest{
				Name:        "asd",
				Description: "asd",
				Permission:  []permissions.Permission{"asd"},
			},
		},
	)
}

func ExampleService_Create() {
	client, _ := misskey.NewClientWithOptions(
		misskey.WithSimpleConfig("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN")),
	)

	resp, err := client.App().Create(app.CreateRequest{
		Name:        "test app",
		Description: "my test app",
		Permission: []permissions.Permission{
			permissions.Write(permissions.Account),
			permissions.Read(permissions.Account),
		},
	})
	if err != nil {
		log.Printf("[Apps/Create] %s", err)

		return
	}

	log.Printf("[Apps/Create] %s created", resp.Name)
}
