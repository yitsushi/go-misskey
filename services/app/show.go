package app

import (
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
)

// ShowRequest represents an Show request.
type ShowRequest struct {
	AppID string `json:"appId"`
}

// Validate the request.
func (r ShowRequest) Validate() error {
	if r.AppID == "" {
		return core.RequestValidationError{
			Request: r,
			Message: core.UndefinedRequiredField,
			Field:   "AppID",
		}
	}

	return nil
}

// Show app.
func (s *Service) Show(appID string) (models.App, error) {
	var response models.App
	err := s.Call(
		&core.JSONRequest{Request: &ShowRequest{AppID: appID}, Path: "/app/show"},
		&response,
	)

	return response, err
}
