package dogo

import (
	"fmt"
	"testing"
)

func TestURLCreation(t *testing.T) {
	table := map[string]string{
		createURL(DropletsEndpoint, nil, nil): fmt.Sprintf("%s/?", DropletsEndpoint),
		createURL(DropletsEndpoint, 123, nil): fmt.Sprintf("%s/%d/?", DropletsEndpoint, 123),
		createURL(DropletsEndpoint, "happy", nil): fmt.Sprintf("%s/%s/?", DropletsEndpoint, "happy"),
	}

	for got, want := range table {
		if got != want {
			t.Errorf("Expected %v, got %v", want, got)
		}
	}
}
