package username_test

import (
	"log"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/services/username"
	"github.com/yitsushi/go-misskey/test"
)

func TestService_Available(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/username/available",
		RequestData:  &username.AvailableRequest{},
		ResponseFile: "available.json",
		StatusCode:   http.StatusOK,
	})

	available, err := client.Username().Available("t")
	if !assert.NoError(t, err) {
		return
	}

	assert.True(t, available, "Username should be available")
}

func TestService_NotAvailable(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/username/available",
		RequestData:  &username.AvailableRequest{},
		ResponseFile: "not-available.json",
		StatusCode:   http.StatusOK,
	})

	available, err := client.Username().Available("t")
	if !assert.NoError(t, err) {
		return
	}

	assert.False(t, available, "Username should not be available")
}

func TestValidate(t *testing.T) {
	tests := []struct {
		name        string
		username    string
		expectError bool
	}{
		{
			name:        "valid",
			username:    "support",
			expectError: false,
		},
		{
			name:        "invalid",
			username:    "bad-username",
			expectError: true,
		},
		{
			name:        "invalid",
			username:    "space containing",
			expectError: true,
		},
		{
			name:        "empty",
			username:    "",
			expectError: true,
		},
		{
			name:        "too long",
			username:    "asdffdsaasdffdsaasdffdsa",
			expectError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := username.AvailableRequest{
				Username: tt.username,
			}.Validate()
			switch tt.expectError {
			case true:
				assert.Error(t, err)
			case false:
				assert.NoError(t, err)
			}
		})
	}
}

func ExampleService_Available() {
	client, _ := misskey.NewClientWithOptions(misskey.WithBaseURL("https", "slippy.xyz", ""))
	username := "t"
	available, err := client.Username().Available(username)

	switch {
	case err != nil:
		log.Printf("[Username/Available] %s", err)
	case available:
		log.Printf("The username %s is available.", username)
	default:
		log.Printf("The username %s is not available.", username)
	}
}
