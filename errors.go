package digitalocean

import (
	"errors"
	"fmt"
)

var (
	EnvError = errors.New("DIGITALOCEAN_TOKEN not found in environment, set the variable.")
)

type APIError struct {
	StatusCode int
	ID         string `json:"id,omitempty"`
	Message    string `json:"message,omitempty"`
}

func (e *APIError) Error() string {
	return fmt.Sprintf("%d %s: %s", e.StatusCode, e.ID, e.Message)
}
