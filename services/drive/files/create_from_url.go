package files

import (
	"context"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// RequestTimout is the timeout of a request in seconds.
const RequestTimout = 30

// CreateFromURLOptions has all the values you can play with.
type CreateFromURLOptions struct {
	FolderID       string
	Name           string
	IsSensitive    bool
	Force          bool
	URL            string
	DownloadClient core.HTTPClient
}

// CreateFromURL downloads a file and then uploads it to the server.
//
// It's not a Misskey endpoint. It's a pseudo endpoint.
//
// Purpose: Replace UploadFromURL with a more convenient approach,
// so we can upload files from external URLs without using the async
// UploadFromURL. That way, we can get back a File from the request.
// And of course, because something stinks to me around that endpoint.
// Why? Check the comment on UploadFromURL.
func (s *Service) CreateFromURL(options CreateFromURLOptions) (models.File, error) {
	return s.Create(CreateRequest{
		FolderID:    options.FolderID,
		Name:        options.Name,
		IsSensitive: options.IsSensitive,
		Force:       options.Force,
		Content:     downloadFile(options.URL, options.DownloadClient),
	})
}

func downloadFile(url string, client core.HTTPClient) []byte {
	req, err := http.NewRequestWithContext(context.Background(), "GET", url, nil)
	if err != nil {
		return []byte{}
	}

	if client == nil {
		client = &http.Client{
			Timeout: time.Second * RequestTimout,
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		return []byte{}
	}

	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}
	}

	return content
}
