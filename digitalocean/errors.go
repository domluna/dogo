package digitalocean

import (
	"fmt"
)

type Error struct {
	StatusCode int
	ID         string `json:"id,omitempty"`
	Message    string `json:"message,omitempty"`
}

func (e Error) Error() string {
	return fmt.Sprintf("%d: %s -- %s\n", e.StatusCode, e.ID, e.Message)
}
