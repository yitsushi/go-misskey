package instance_test

import (
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/services/admin/instance"
	"github.com/yitsushi/go-misskey/test"
)

func TestService_ServerInfo(t *testing.T) {
	client := test.MakeMockClient(test.SimpleMockOptions{
		Endpoint:     "/api/admin/server-info",
		RequestData:  &instance.ServerInfoRequest{},
		ResponseFile: "server-info.json",
		StatusCode:   http.StatusOK,
	})

	response, err := client.Admin().Instance().ServerInfo()

	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, "string", response.Machine, "machine")
	assert.Equal(t, "linux", response.OS, "os")
	assert.Equal(t, "20.2.0", response.Node, "node")
	assert.Equal(t, "string", response.PSQL, "psql")
	assert.EqualValues(t, 20, response.CPU.Cores, "cpu.cores")
	assert.EqualValues(t, "arm64", response.CPU.Model, "cpu.model")
	assert.EqualValues(t, 5254, response.Mem.Total, "mem.total")
	assert.EqualValues(t, 654, response.FS.Total, "fs.total")
	assert.EqualValues(t, 3, response.FS.Used, "fs.used")
	assert.EqualValues(t, "eth0", response.Net.Interface, "net.interface")
}

// ExampleService_ServerInfo demonstrates how to use Admin.Instance.ServerInfo.
func ExampleService_ServerInfo() {
	client, _ := misskey.NewClientWithOptions(misskey.WithSimpleConfig("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN")))

	response, err := client.Admin().Instance().ServerInfo()
	if err != nil {
		log.Printf("[Admin/Instance/ServerInfo] %s", err)

		return
	}

	log.Printf("Server Info: %v", response)
}
