package users_test

import (
	"log"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/services/users"
	"github.com/yitsushi/go-misskey/test"
)

func TestService_Available(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/username/available",
		RequestData:  &users.AvailableRequest{},
		ResponseFile: "username-available.json",
		StatusCode:   http.StatusOK,
	})

	available, err := client.Users().IsUsernameAvailable("t")
	if !assert.NoError(t, err) {
		return
	}

	assert.True(t, available, "Username should be available")
}

func TestService_NotAvailable(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/username/available",
		RequestData:  &users.AvailableRequest{},
		ResponseFile: "username-not-available.json",
		StatusCode:   http.StatusOK,
	})

	available, err := client.Users().IsUsernameAvailable("support")
	if !assert.NoError(t, err) {
		return
	}

	assert.False(t, available, "Username should not be available")
}

func TestValidate(t *testing.T) {
	err := users.AvailableRequest{
		Username: "support",
	}.Validate()
	assert.NoError(t, err)
}

func TestValidate_Invalid(t *testing.T) {
	tests := []struct {
		name     string
		username string
	}{
		{
			name:     "invalid",
			username: "bad-username",
		},
		{
			name:     "invalid",
			username: "space containing",
		},
		{
			name:     "empty",
			username: "",
		},
		{
			name:     "too long",
			username: "asdffdsaasdffdsaasdffdsa",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := users.AvailableRequest{
				Username: tt.username,
			}.Validate()
			assert.Error(t, err)
		})
	}
}

func ExampleService_IsUsernameAvailable() {
	client, _ := misskey.NewClientWithOptions(misskey.WithBaseURL("https", "slippy.xyz", ""))
	username := "admin"
	available, err := client.Users().IsUsernameAvailable(username)

	switch {
	case err != nil:
		log.Printf("[Username/Available] %s", err)
	case available:
		log.Printf("The username %s is available.", username)
	default:
		log.Printf("The username %s is not available.", username)
	}
}
