package models

import "github.com/yitsushi/go-misskey/core"

// App represents an App from /app endpoints.
type App struct {
	ID           string      `json:"id"`
	Name         string      `json:"name"`
	CallbackURL  core.String `json:"callbackUrl"`
	Permission   []string    `json:"permission"`
	Secret       string      `json:"secret"`
	IsAuthorized bool        `json:"isAuthorized"`
}
