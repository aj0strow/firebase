package firebase

import (
	"net/http"
	"time"
)

type App struct {
	DatabaseURL string
	Secret      string
	Timeout     time.Duration
	Transport   http.RoundTripper
}

func (app *App) Auth(auth string) *Client {
	return &Client{
		client: &http.Client{
			Timeout:   app.Timeout,
			Transport: app.Transport,
		},
		databaseURL: app.DatabaseURL,
		auth:        auth,
	}
}

func (app *App) Admin() *Client {
	return app.Auth(app.Secret)
}
