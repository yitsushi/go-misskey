package files

import (
	"context"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/yitsushi/go-misskey/models"
)

// CreateFromURLOptions has all the values you can play with.
type CreateFromURLOptions struct {
	FolderID    string
	Name        string
	IsSensitive bool
	Force       bool
	URL         string
}

// CreateFromURL downloads a file and then uploads to the server.
//
// It's not a Misskey endpoint. It's a pseudo endpoint.
//
// Purpose: Replace UploadFromURL with a more convinient approach,
// so we can upload files from external URLs without using the async
// UploadFromURL. That way, we can get back a File from the request.
// And of course, because something stinks to me around that endpoint.
// Why? Check the comment on UploadFromURL.
func (s *Service) CreateFromURL(options *CreateFromURLOptions) (models.File, error) {
	return s.Create(&CreateOptions{
		FolderID:    options.FolderID,
		Name:        options.Name,
		IsSensitive: options.IsSensitive,
		Force:       options.Force,
		Content:     downloadFile(options.URL),
	})
}

func downloadFile(url string) []byte {
	req, err := http.NewRequestWithContext(context.Background(), "GET", url, nil)
	if err != nil {
		return []byte{}
	}

	client := &http.Client{
		Timeout: time.Minute,
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
