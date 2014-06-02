package dogo

import (
	"testing"
)

func TestSSHLogin(t *testing.T) {
	err := Login("root", "107.170.186.65", 22)
	if err != nil {
		t.Errorf("Expected %v, got %v", nil, err)
	}
}
