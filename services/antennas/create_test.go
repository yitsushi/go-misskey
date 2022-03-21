package antennas_test

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
	"github.com/yitsushi/go-misskey/services/antennas"
	"github.com/yitsushi/go-misskey/test"
)

func TestService_Create(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/antennas/create",
		RequestData:  &antennas.CreateRequest{},
		ResponseFile: "create.json",
		StatusCode:   http.StatusOK,
	})

	antenna, err := client.Antennas().Create(antennas.CreateRequest{
		Name:   "Test name",
		Source: models.AllSrc,
		Keywords: [][]string{
			{"test"},
		},
		ExcludeKeywords: [][]string{},
		Users:           []string{},
	})
	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, "Test name", antenna.Name)
}

func TestService_Create_missingField(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/antennas/create",
		RequestData:  &antennas.CreateRequest{},
		ResponseFile: "create.json",
		StatusCode:   http.StatusOK,
	})

	_, err := client.Antennas().Create(antennas.CreateRequest{})
	if !assert.Error(t, err) {
		return
	}

	assert.Contains(t, err.Error(), "[Name]")
}

func TestCreateRequest_Validate(t *testing.T) {
	test.ValidateRequests(
		t,
		[]core.BaseRequest{
			antennas.CreateRequest{},
			antennas.CreateRequest{
				Name: "This is a name",
			},
			antennas.CreateRequest{
				Name:   "This is a name",
				Source: models.AllSrc,
			},
			antennas.CreateRequest{
				Name:   "This is a name",
				Source: models.AllSrc,
				Keywords: [][]string{
					{"test"},
				},
			},
			antennas.CreateRequest{
				Name:   "This is a name",
				Source: models.AllSrc,
				Keywords: [][]string{
					{"test"},
				},
				ExcludeKeywords: [][]string{},
			},
		},
		[]core.BaseRequest{
			antennas.CreateRequest{
				Name:   "This is a name",
				Source: models.AllSrc,
				Keywords: [][]string{
					{"test"},
				},
				ExcludeKeywords: [][]string{},
				Users:           []string{},
			},
		},
	)
}

func ExampleService_Create() {
	client, _ := misskey.NewClientWithOptions(misskey.WithSimpleConfig("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN")))

	resp, err := client.Antennas().Create(antennas.CreateRequest{
		Name:            "test",
		Source:          models.AllSrc,
		UserListID:      nil,
		UserGroupID:     nil,
		Keywords:        [][]string{{"update", "what"}, {"stuff"}},
		ExcludeKeywords: [][]string{},
		Users:           []string{},
		CaseSensitive:   false,
		WithReplies:     true,
		WithOnlyFile:    true,
		Notify:          false,
	})
	if err != nil {
		log.Printf("[Antennas/Create] %s", err)

		return
	}

	log.Printf("[Antennas/Create] %s created", resp.Name)
}
