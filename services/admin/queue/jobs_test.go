package queue_test

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/services/admin/queue"
	"github.com/yitsushi/go-misskey/test"
)

func jobServer() *misskey.Client {
	return test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:    "/api/admin/queue/jobs",
		RequestData: &queue.JobsRequest{},
		ResponseFileFunc: func(r interface{}) string {
			request := r.(*queue.JobsRequest)

			switch request.Domain {
			case queue.DeliverDomain:
				return "jobs-deliver.json"
			case queue.InboxDomain:
				return "jobs-inbox.json"
			case queue.DBDomain, queue.ObjectStorageDomain:
				return ""
			}

			return ""
		},
		StatusCode: http.StatusOK,
	})
}

func TestService_Jobs_deliver(t *testing.T) {
	client := jobServer()

	response, err := client.Admin().Queue().Jobs(queue.JobsRequest{
		Domain: queue.DeliverDomain,
		State:  queue.DelayedState,
	})

	if !assert.NoError(t, err) {
		return
	}

	assert.Len(t, response, 2)
}

func TestService_Jobs_inbox(t *testing.T) {
	client := jobServer()

	response, err := client.Admin().Queue().Jobs(queue.JobsRequest{
		Domain: queue.InboxDomain,
		State:  queue.DelayedState,
	})

	if !assert.NoError(t, err) {
		return
	}

	assert.Len(t, response, 1)
}

func TestJobsRequest_Validate(t *testing.T) {
	test.ValidateRequests(
		t,
		[]core.BaseRequest{
			queue.JobsRequest{},
			queue.JobsRequest{Domain: queue.DeliverDomain},
			queue.JobsRequest{State: queue.ActiveState},
		},
		[]core.BaseRequest{
			queue.JobsRequest{
				Domain: queue.DeliverDomain,
				State:  queue.ActiveState,
			},
		},
	)
}

func ExampleService_Jobs() {
	client, _ := misskey.NewClientWithOptions(misskey.WithSimpleConfig("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN")))

	response, err := client.Admin().Queue().Jobs(queue.JobsRequest{})
	if err != nil {
		log.Printf("[Admin/Queue] %s", err)

		return
	}

	for _, item := range response {
		log.Printf("[Admin/Queue] Attempts: %d/%d", item.Attempts, item.MaxAttempts)
	}
}
