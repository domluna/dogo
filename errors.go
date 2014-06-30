package dogo

import (
// "fmt"
)

type Error struct {
	Status int
	Type   string `json:"error,omitempty"`
	Desc   string `json:"description,omitempty"`
}

func (e Error) Error() string {
	return ""
}
