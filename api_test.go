package dogo

import (
	// "encoding/json"
	// "io/ioutil"
	// "net/http"
	// "net/http/httptest"
	"os"
	"testing"
)

const (
	token = "your token here"
)

func TestEnvAuth(t *testing.T) {
	msg := "DIGITALOCEAN_TOKEN not found in environment"

	cli, err := EnvAuth()
	if err.Error() != msg {
		t.Errorf("Expected %v, got %v", msg, err)
	}

	os.Setenv("DIGITALOCEAN_TOKEN", "mytokenhere")
	temp := Client{"mytokenhere"}
	cli, err = EnvAuth()
	if cli != temp {
		t.Errorf("Expected %v, got %v", temp, cli)
	}
}

func TestGet(t *testing.T) {
	t.Logf("Starting Get tests ... \n")
	c := Client{token}
	dc := DropletClient{c}
	_, err := dc.GetAll()
	if err != nil {
		t.Fatal(err)
	}

	sc := SizeClient{c}
	_, err = sc.GetAll()
	if err != nil {
		t.Fatal(err)
	}

	rc := RegionClient{c}
	_, err = rc.GetAll()
	if err != nil {
		t.Fatal(err)
	}

	kc := KeyClient{c}
	_, err = kc.GetAll()
	if err != nil {
		t.Fatal(err)
	}

	doc := DomainClient{c}
	_, err = doc.GetAll()
	if err != nil {
		t.Fatal(err)
	}
}

func TestCreateDroplet(t *testing.T) {
	t.Logf("Starting Create Test ... \n")
	c := Client{token}
	dc := DropletClient{c}
	droplet, err := dc.Create(map[string]interface{}{
		"name":   "test-create",
		"region": "nyc2",
		"size":   "512mb",
		"image":  4296335,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%v\n", droplet)
}

func TestPut(t *testing.T) {

}

func TestDelete(t *testing.T) {
	t.Logf("Starting Delete Test ... \n")
	c := Client{token}
	dc := DropletClient{c}
	err := dc.Destroy(1975885)
	if err != nil {
		t.Fatal(err)
	}
}
