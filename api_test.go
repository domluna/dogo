package dogo

import (
	// "encoding/json"
	"fmt"
	// "io/ioutil"
	// "net/http"
	// "net/http/httptest"
	"os"
	"testing"
)

const (
	token = "bcae7169e90c0df91d95feaf61bdc0126fef129ca4b47bd8d2e8b6bdf54bfb98"
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
	c := Client{token}
	dc := DropletClient{c}
	d, err := dc.GetDroplet(1618748)
	// d, err := dc.GetDroplets()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Droplets: %v\n", d)

	sc := SizeClient{c}
	s, err := sc.GetSizes()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Sizes: %v\n", s)
}

func TestPost(t *testing.T) {

}

func TestPut(t *testing.T) {

}

func TestDel(t *testing.T) {

}
