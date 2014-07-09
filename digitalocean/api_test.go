package digitalocean

import (
	// "encoding/json"
	// "io/ioutil"
	// "net/http"
	// "net/http/httptest"
	"os"
	"testing"
)

func TestEnvAuth(t *testing.T) {
	os.Setenv("DIGITALOCEAN_TOKEN", "")
	token, err := EnvAuth()
	if err != EnvError {
		t.Errorf("Expected %v, got %v", EnvError, err)
	}

	os.Setenv("DIGITALOCEAN_TOKEN", "mytokenhere")
	temp := "mytokenhere"
	token, err = EnvAuth()
	if token != temp {
		t.Errorf("Expected %v, got %v", temp, token)
	}
}
