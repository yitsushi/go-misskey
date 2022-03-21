package notes_test

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
	"github.com/yitsushi/go-misskey/services/notes"
	"github.com/yitsushi/go-misskey/test"
)

func TestService_Create(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/notes/create",
		RequestData:  &notes.CreateRequest{},
		ResponseFile: "create.json",
		StatusCode:   http.StatusOK,
	})

	response, err := client.Notes().Create(notes.CreateRequest{
		Text:       core.NewString("test"),
		Visibility: models.VisibilityHome,
		Poll: &notes.Poll{
			Choices: []string{"a", "b", "c"},
		},
	})
	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, "test", response.CreatedNote.Text)
}

func TestCreateRequest_Validate(t *testing.T) {
	test.ValidateRequests(
		t,
		[]core.BaseRequest{
			notes.CreateRequest{},
			notes.CreateRequest{
				Text: core.NewString(""),
			},
			notes.CreateRequest{
				Poll: &notes.Poll{},
			},
			notes.CreateRequest{
				Poll: &notes.Poll{
					Choices: []string{"a"},
				},
			},
		},
		[]core.BaseRequest{
			notes.CreateRequest{
				Text: core.NewString("test"),
			},
			notes.CreateRequest{
				FileIDs: []string{"file"},
			},
			notes.CreateRequest{
				Poll: &notes.Poll{
					Choices: []string{"a", "b"},
				},
			},
		},
	)
}

func ExampleService_Create() {
	client, _ := misskey.NewClientWithOptions(misskey.WithSimpleConfig("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN")))
	client.LogLevel(logrus.DebugLevel)

	response, err := client.Notes().Create(notes.CreateRequest{
		Text:       core.NewString("test"),
		Visibility: models.VisibilityHome,
		Poll: &notes.Poll{
			Choices: []string{"a", "b", "c"},
		},
	})
	if err != nil {
		log.Printf("[Notes] Error happened: %s", err)

		return
	}

	log.Println(response.CreatedNote.ID)
}
