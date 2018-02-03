package firebase

import (
	"net/http"
	"time"
)

// App keeps track of settings to initialize clients.
type App struct {
	DatabaseURL string
	Secret      string
	Timeout     time.Duration
	Transport   http.RoundTripper
}

// Auth initializes a new client with the passed in auth.
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

// Admin initalizes a new client with the secret key.
func (app *App) Admin() *Client {
	return app.Auth(app.Secret)
}
