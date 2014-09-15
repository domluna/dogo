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

func makeClient(t *testing.T) *Client {
	client, err := NewClient("foo")
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	if client.Token != "foo" {
		t.Fatalf("Expected Token to be %v, got %v", "foo", client.Token)
	}
	return client
}

func Test_NewClient_EnvPresent(t *testing.T) {
	os.Setenv("DIGITALOCEAN_TOKEN", "foo")
	client, err := NewClient("")

	if err != nil {
		t.Errorf("%s", err)
	}

	if client.Token != "foo" {
		t.Errorf("Expected Token to be %v, got %v", "foo", client.Token)
	}

	os.Setenv("DIGITALOCEAN_TOKEN", "foo")
	client, err = NewClient("")
}

func Test_NewClient_EnvNotPresent(t *testing.T) {
	os.Setenv("DIGITALOCEAN_TOKEN", "")
	_, err := NewClient("")

	if err == nil {
		t.Errorf("Should be an error about DIGITALOCEAN_TOKEN not being present")
	}

	if err != EnvError {
		t.Errorf("Expected %s, got %s", EnvError, err)
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
		Droplets `json:"droplets"`
	}{}

	ts := httptest.NewServer(http.HandlerFunc(f))
	defer ts.Close()

	client := makeClient(t)
	client.URL = ts.URL
	err := client.get("", &s)
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
		Droplet `json:"droplet"`
	}{}

	params := Params{
		"name":   "joker",
		"size":   "512mb",
		"region": "gotham",
		"image":  42,
	}

	client := makeClient(t)
	client.URL = ts.URL
	err := client.post("", params, &s)
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

	client := makeClient(t)
	client.URL = ts.URL
	err := client.delete("")
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

	client := makeClient(t)
	client.URL = ts.URL
	err := client.delete("")
	errMsg := "501 prison escape: this is gotham we're talking about!"
	if err == nil {
		t.Error("expected there to be an error")
	}

	if err.Error() != errMsg {
		t.Errorf("expected %v, got %v", errMsg, err.Error())
	}
}
