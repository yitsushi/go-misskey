package users_test

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/services/admin/users"
	"github.com/yitsushi/go-misskey/test"
)

func TestService_ResetPassword(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/admin/reset-password",
		RequestData:  &users.ResetPasswordRequest{},
		ResponseFile: "reset-password.json",
		StatusCode:   http.StatusOK,
	})

	newPassword, err := client.Admin().Users().ResetPassword("83sv4lyx22")
	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, "jashdkjahsjdk", newPassword)
}

func TestResetPasswordRequest_Validate(t *testing.T) {
	test.ValidateRequests(
		t,
		[]core.BaseRequest{
			users.ResetPasswordRequest{},
		},
		[]core.BaseRequest{
			users.ResetPasswordRequest{UserID: "asd"},
		},
	)
}

func ExampleService_ResetPassword() {
	client, _ := misskey.NewClientWithOptions(misskey.WithSimpleConfig("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN")))

	newPassword, err := client.Admin().Users().ResetPassword("83sv4lyx22")
	if err != nil {
		log.Printf("[Admin/Users] %s", err)

		return
	}

	log.Printf("[Admin/Users] New password: %s", newPassword)
}
