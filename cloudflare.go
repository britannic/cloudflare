package main

import (
	"os"

	cloudflare "github.com/cloudflare/cloudflare-go"
)

// See Cloudflare instructions here: https://api.cloudflare.com/#getting-started-endpoints
func (o *opts) getCFAPI() (*cloudflare.API, error) {
	var (
		email = os.Getenv("CF_API_EMAIL")
		key   = os.Getenv("CF_API_KEY")
	)

	if *o.apiKey != "" {
		key = *o.apiKey
	}

	if *o.email != "" {
		email = *o.email
	}

	return cloudflare.New(key, email)
}
