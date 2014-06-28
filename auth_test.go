package dogo

import (
	"os"
	"testing"
)

func TestEnvAuth(t *testing.T) {
	msg := "DIGITALOCEAN_TOKEN not found in environment"

	cli, err := EnvAuth()
	if err.Error() != msg {
		t.Errorf("Expected %v, got %v", msg, err)
	}

	os.Setenv("DIGITALOCEAN_TOKEN", "mytokenhere")
	temp := Client{"mytokenhere"}
	if cli != temp {
		t.Errorf("Expected %v, got %v", temp, cli)
	}

}
