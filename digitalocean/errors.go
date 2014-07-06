package digitalocean

import (
	"errors"
	"fmt"
)

var (
	EnvError = errors.New("DIGITALOCEAN_TOKEN not found in environment")
)

type APIError struct {
	StatusCode int
	ID         string `json:"id,omitempty"`
	Message    string `json:"message,omitempty"`
}

func (e *APIError) Error() string {
	return fmt.Sprintf("%d %s: %s\n", e.StatusCode, e.ID, e.Message)
}
