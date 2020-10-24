package files

import (
	"github.com/yitsushi/go-misskey/core"
)

// UploadFromURLRequest is the request to upload a file from a URL.
type UploadFromURLRequest struct {
	URL         string      `json:"url"`
	Name        string      `json:"name"`
	Comment     core.String `json:"comment"`
	Marker      core.String `json:"marker"`
	Force       bool        `json:"force"`
	IsSensitive bool        `json:"isSensitive"`
}

// Validate the request.
func (r UploadFromURLRequest) Validate() error {
	return nil
}

// UploadFromURL asks the server to download an image from an external URL.
// As it's an asnyc request, we know nothing about the file yet, so
// the only response is error or nothing.
//
// It seems the server completely ignores the Name, Comment and Marker
// fields. To be honest, I don't even know what Comment and Marker means,
// as there is no "alt text" support on images with Misskey. I think it was
// supported with v11 and it was never re-implemented in v12.
//
// Advise: Use CreateFromURL instead of UploadFromURL.
func (s *Service) UploadFromURL(request UploadFromURLRequest) error {
	return s.Call(
		&core.JSONRequest{Request: &request, Path: "/drive/files/upload-from-url"},
		&core.DummyResponse{},
	)
}
