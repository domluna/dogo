package dogo

import (
	"os"
	"testing"
)

func TestEnvAuth(t *testing.T) {
	k := "DIGITALOCEAN_API_KEY not found in environment"
	c := "DIGITALOCEAN_CLIENT_ID not found in environment"

	auth, err := EnvAuth()
	if err.Error() != k {
		t.Errorf("Expected %v, got %v", k, err)
	}

	os.Setenv("DIGITALOCEAN_API_KEY", "awesomekey")
	auth, err = EnvAuth()
	if err.Error() != c {
		t.Errorf("Expected %v, got %v", c, err)
	}

	os.Setenv("DIGITALOCEAN_CLIENT_ID", "awesomeclient")
	auth, err = EnvAuth()
	if err != nil {
		t.Errorf("Expected %v, got %v", nil, err)
	}

	auth2 := Auth{"awesomeclient", "awesomekey"}
	if auth != auth2 {
		t.Errorf("Expected %v, got %v", auth2, auth)
	}

}
