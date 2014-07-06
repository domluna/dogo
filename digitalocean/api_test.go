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

func TestGet(t *testing.T) {
	t.Logf("Starting Get tests ... \n")
}

// func TestCreateDroplet(t *testing.T) {
// 	t.Logf("Starting Create Test ... \n")
// 	c := Client{token}
// 	dc := DropletClient{c}
// 	droplet, err := dc.Create(map[string]interface{}{
// 		"name":   "test-create",
// 		"region": "nyc2",
// 		"size":   "512mb",
// 		"image":  4296335,
// 	})
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	t.Logf("%v\n", droplet)
// }
//
// func TestPut(t *testing.T) {
//
// }
//
// func TestDelete(t *testing.T) {
// 	t.Logf("Starting Delete Test ... \n")
// 	c := Client{token}
// 	dc := DropletClient{c}
// 	err := dc.Destroy(1975885)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// }
