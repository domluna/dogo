package digitalocean

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

func Test_ListDomains(t *testing.T) {
	setup(t)
	defer teardown()

	want := &Domain{
		Name:     "example.com",
		TTL:      1800,
		ZoneFile: "Example zone file text...",
	}

	mux.HandleFunc("/domains", func(w http.ResponseWriter, r *http.Request) {
		assertEqual(t, r.Method, "GET")
		fmt.Fprint(w, listDomainsExample)
	})

	domains, err := client.ListDomains()

	assertEqual(t, err, nil)
	assertEqual(t, len(domains), 1)
	assertEqual(t, domains[0].Name, want.Name)
	assertEqual(t, domains[0].TTL, want.TTL)
	assertEqual(t, domains[0].ZoneFile, want.ZoneFile)
}

func Test_GetDomain(t *testing.T) {
	setup(t)
	defer teardown()

	want := &Domain{
		Name:     "example.com",
		TTL:      1800,
		ZoneFile: "Example zone file text...",
	}

	mux.HandleFunc("/domains/example.com", func(w http.ResponseWriter, r *http.Request) {
		assertEqual(t, r.Method, "GET")
		fmt.Fprint(w, domainExample)
	})

	domain, err := client.GetDomain(want.Name)
	assertEqual(t, err, nil)
	assertEqual(t, domain.Name, want.Name)
	assertEqual(t, domain.TTL, want.TTL)
	assertEqual(t, domain.ZoneFile, want.ZoneFile)
}

func Test_CreateDomain(t *testing.T) {
	setup(t)
	defer teardown()

	want := &Domain{
		Name:     "example.com",
		TTL:      1800,
		ZoneFile: "Example zone file text...",
	}

	opts := &CreateDomainOpts{
		Name:      "example.com",
		IPAddress: "127.0.0.20",
	}

	mux.HandleFunc("/domains", func(w http.ResponseWriter, r *http.Request) {
		v := new(CreateDomainOpts)
		json.NewDecoder(r.Body).Decode(v)
		assertEqual(t, r.Method, "POST")
		assertEqual(t, v, opts)
		fmt.Fprint(w, domainExample)
	})

	domain, err := client.CreateDomain(opts)
	assertEqual(t, err, nil)
	assertEqual(t, domain.Name, want.Name)
	assertEqual(t, domain.TTL, want.TTL)
	assertEqual(t, domain.ZoneFile, want.ZoneFile)
}

func Test_DeleteDomain(t *testing.T) {
	setup(t)
	defer teardown()

	name := "example.com"
	mux.HandleFunc("/domains/example.com", func(w http.ResponseWriter, r *http.Request) {
		assertEqual(t, r.Method, "DELETE")
		fmt.Fprint(w, "")
	})

	err := client.DeleteDomain(name)
	assertEqual(t, err, nil)
}

var listDomainsExample = `{ 
  "domains": [
    {
      "name": "example.com",
      "ttl": 1800,
      "zone_file": "Example zone file text..."
    }
  ],
  "meta": {
    "total": 1
  }
}`

var domainExample = `{
  "domain": {
    "name": "example.com",
    "ttl": 1800,
    "zone_file": "Example zone file text..."
  }
}`
