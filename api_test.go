package digitalocean

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
)

var (
	// mux is the HTTP request multiplexer used with the test server.
	mux *http.ServeMux

	// http server used to mock api responses.
	server *httptest.Server

	// DigitalOcean client to be tested.
	client *Client
)

// Sets up server and client on which the tests will
// be run.
func setup(t *testing.T) {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
	client = makeClient(t)
	client.URL = server.URL
}

// Closes the http server when tests are done.
func teardown() {
	server.Close()
}

func writeJSON(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf8")
	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		panic("writeJSON: " + err.Error())
	}
}

func assertEqual(t *testing.T, got interface{}, want interface{}) {
	if !reflect.DeepEqual(want, got) {
		t.Errorf("Expected %v, got %v", want, got)
	}
}

func assertNotEqual(t *testing.T, got interface{}, want interface{}) {
	if reflect.DeepEqual(want, got) {
		t.Errorf("Expected %v, got %v", want, got)
	}
}

// utility to create a test server with sane defaults for faster testing
func testServer(status int, output string) {
	f := func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(status)
		fmt.Fprintln(w, output)
	}
	mux.HandleFunc("/", f)
}

// Utility function for creating a Client.
// Also checks to make sure the functionality
// is working as expected.
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
	assertEqual(t, err, EnvError)
}

func Test_API_Get(t *testing.T) {
	setup(t)
	defer teardown()
	f := func(w http.ResponseWriter, req *http.Request) {
		assertEqual(t, req.Method, "GET")
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

	mux.HandleFunc("/", f)
	err := client.get("", &s)
	name := s.Droplets[0].Name
	assertEqual(t, err, nil)
	assertEqual(t, len(s.Droplets), 2)
	assertEqual(t, name, "batman")
}

func Test_API_Post(t *testing.T) {
	setup(t)
	defer teardown()
	f := func(w http.ResponseWriter, req *http.Request) {
		assertEqual(t, req.Header.Get("Content-Type"), "application/json")
		assertEqual(t, req.Method, "POST")

		body, _ := ioutil.ReadAll(req.Body)
		var data Params
		json.Unmarshal(body, &data)

		assertEqual(t, data["size"], "512mb")
		assertEqual(t, data["name"], "joker")
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

	mux.HandleFunc("/", f)
	s := struct {
		Droplet `json:"droplet"`
	}{}
	params := Params{
		"name":   "joker",
		"size":   "512mb",
		"region": "gotham",
		"image":  42,
	}

	err := client.post("", params, &s)
	name := s.Droplet.Name
	assertEqual(t, name, "joker")
	assertEqual(t, err, nil)
}

func Test_API_Put(t *testing.T) {
	setup(t)
	defer teardown()
	f := func(w http.ResponseWriter, req *http.Request) {
		assertEqual(t, req.Header.Get("Content-Type"), "application/x-www-form-urlencoded")
		assertEqual(t, req.Method, "PUT")

		body, _ := ioutil.ReadAll(req.Body)
		var data Params
		json.Unmarshal(body, &data)

		assertEqual(t, data["name"], "joker")
		assertEqual(t, data["size"], "512mb")
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

	mux.HandleFunc("/", f)
	s := struct {
		Droplet `json:"droplet"`
	}{}
	params := Params{
		"name":   "joker",
		"size":   "512mb",
		"region": "gotham",
		"image":  42,
	}

	err := client.put("", params, &s)
	name := s.Droplet.Name
	assertEqual(t, err, nil)
	assertEqual(t, name, "joker")

}

func Test_API_Delete(t *testing.T) {
	setup(t)
	defer teardown()
	f := func(w http.ResponseWriter, req *http.Request) {
		assertEqual(t, req.Method, "DELETE")
		w.WriteHeader(204)
	}

	mux.HandleFunc("/", f)
	err := client.delete("")
	assertEqual(t, err, nil)
}

func Test_API_Errors(t *testing.T) {
	setup(t)
	defer teardown()
	f := func(w http.ResponseWriter, req *http.Request) {
		assertEqual(t, req.Method, "DELETE")
		w.WriteHeader(501)
		fmt.Fprintln(w, `{
			"id": "prison escape",
			"message": "this is gotham we're talking about!"
		}`)
	}

	mux.HandleFunc("/", f)
	err := client.delete("")
	errMsg := "501 prison escape: this is gotham we're talking about!"
	assertNotEqual(t, err, nil)
	assertEqual(t, err.Error(), errMsg)
}
