package dogo

import (
	"errors"
	"fmt"
)

var (
	// ErrEnv is a error for the DigitalOcean environment token being unset.
	ErrEnv = errors.New("DIGITALOCEAN_TOKEN not found in environment, set the variable")
)

// APIError is an error wrapping the HTTP response back from
// the DigitalOcean servers.
type APIError struct {
	StatusCode int
	ID         string `json:"id,omitempty"`
	Message    string `json:"message,omitempty"`
}

func (e *APIError) Error() string {
	return fmt.Sprintf("%d %s: %s", e.StatusCode, e.ID, e.Message)
}
