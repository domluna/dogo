package dogo

import (
	"errors"
	// "net/http"
	"os"
)

// Auth contains data required to authenticate
// get DigitalOcean api.
type Auth struct {
	ClientID string
	APIKey   string
}

// Client is a wrapper around Auth, Clients are used
// to query the api.
//
// To make a new Client call NewClient.
type Client struct {
	Auth
}

// NewClient creates a new Client.
func NewClient(auth Auth) *Client {
	return &Client{auth}
}

// EnvAuth creates an Auth based on environment the
// DIGITALOCEAN_API_KEY and DIGITALOCEAN_CLIENT_ID
// environment variables.
func EnvAuth() (Auth, error) {
	var auth Auth
	auth.APIKey = os.Getenv("DIGITALOCEAN_API_KEY")
	auth.ClientID = os.Getenv("DIGITALOCEAN_CLIENT_ID")
	if auth.APIKey == "" {
		return auth, errors.New("DIGITALOCEAN_API_KEY not found in environment")
	}

	if auth.ClientID == "" {
		return auth, errors.New("DIGITALOCEAN_CLIENT_ID not found in environment")
	}
	return auth, nil
}
