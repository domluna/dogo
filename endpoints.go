package dogo

import (
	"fmt"
)

const (
	// Base URL to API
	BaseURL = "https://api.digitalocean.com/v2"
)

var (
	DropletsEndpoint = endpoint("droplets")
	DomainsEndpoint  = endpoint("domains")
	ImagesEndpoint   = endpoint("images")
	KeysEndpoint     = endpoint("account/keys")
	RegionsEndpoint  = endpoint("regions")
	ActionsEndpoint  = endpoint("actions")
	SizesEndpoint    = endpoint("sizes")
	EventsEndpoint   = endpoint("events")
)

// Create an endpoint by appending resource to base url
func endpoint(resource string) string {
	return fmt.Sprintf("%s/%s", BaseURL, resource)
}
