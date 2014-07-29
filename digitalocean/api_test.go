package digitalocean

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

// Purely for testing purposes, real droplets are way cooler.
type DummyDroplet struct {
	ID     int    `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Region string `json:"region,omitempty"`
	Image  int    `json:"image,omitempty"`
	Size   string `json:"size,omitempty"`
}

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
	f := func(w http.ResponseWriter, req *http.Request) {
		if req.Method != "GET" {
			t.Errorf("expected method to be GET, got %v", req.Method)
		}
		fmt.Fprintln(w, `{
			"droplets": [
				{
					"id": 1,
					"name": "batman"
				},
				{
					"id": 2,
					"name": "robin"
				}
			]
		}`)
	}

	s := struct {
		Droplets []DummyDroplet `json:"droplets"`
	}{}

	ts := httptest.NewServer(http.HandlerFunc(f))
	defer ts.Close()

	client := NewClient("not_actual_token")
	err := client.Get(ts.URL, &s)
	if err != nil {
		t.Error(err)
	}

	if len(s.Droplets) != 2 {
		t.Errorf("Expected 2, got %v (droplets: %v)", len(s.Droplets), s.Droplets)
	}

	name := s.Droplets[0].Name
	if name != "batman" {
		t.Errorf("Expected batman, got %v (droplets: %v)", name, s.Droplets)
	}

}

func TestPost(t *testing.T) {
	f := func(w http.ResponseWriter, req *http.Request) {
		if req.Header.Get("Content-Type") != "application/json" {
			t.Errorf("Expected application/json, got %s", req.Header.Get("Content-Type"))
		}
		if req.Method != "POST" {
			t.Errorf("expected method to be POST, got %v", req.Method)
		}

		body, _ := ioutil.ReadAll(req.Body)
		var data Params
		json.Unmarshal(body, &data)

		if data["name"] != "joker" {
			t.Errorf("Expected joker, got %v", data["name"])
		}

		if data["size"] != "512mb" {
			t.Errorf("Expected 512mb, got %v", data["size"])
		}
		// Respond back
		w.WriteHeader(201)
		fmt.Fprintln(w, `{
			"droplet": {
				"id": 3,
				"name": "joker",
				"region": "gotham",
				"size": "512mb",
				"image": 42
			}
		}`)
	}

	ts := httptest.NewServer(http.HandlerFunc(f))
	defer ts.Close()

	s := struct {
		Droplet DummyDroplet `json:"droplet"`
	}{}

	params := Params{
		"name":   "joker",
		"size":   "512mb",
		"region": "gotham",
		"image":  42,
	}

	client := NewClient("not_actual_token")

	err := client.Post(ts.URL, params, &s)
	if err != nil {
		t.Error(err)
	}

	name := s.Droplet.Name
	if name != "joker" {
		t.Errorf("Expected joker, got %v (droplet: %v)", name, s.Droplet)
	}

}

func TestDelete(t *testing.T) {
	f := func(w http.ResponseWriter, req *http.Request) {
		if req.Method != "DELETE" {
			t.Errorf("expected method to be DELETE, got %v", req.Method)
		}
		fmt.Fprintln(w, "")
	}

	ts := httptest.NewServer(http.HandlerFunc(f))
	defer ts.Close()

	client := NewClient("not_actual_token")
	err := client.Delete(ts.URL)
	if err != nil {
		t.Error(err)
	}
}

func TestErrMessage(t *testing.T) {
	f := func(w http.ResponseWriter, req *http.Request) {
		if req.Method != "DELETE" {
			t.Errorf("expected method to be DELETE, got %v", req.Method)
		}
		w.WriteHeader(501)
		fmt.Fprintln(w, `{
			"id": "prison escape",
			"message": "this is gotham we're talking about!"
		}`)
	}

	ts := httptest.NewServer(http.HandlerFunc(f))
	defer ts.Close()

	client := NewClient("not_actual_token")
	err := client.Delete(ts.URL)
	errMsg := "501 prison escape: this is gotham we're talking about!"
	if err == nil {
		t.Error("expected there to be an error")
	}

	if err.Error() != errMsg {
		t.Errorf("expected %v, got %v", errMsg, err.Error())
	}
}
