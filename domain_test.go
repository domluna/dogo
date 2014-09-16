package digitalocean

import (
	"fmt"
	"testing"
)

func Test_ListDomains(t *testing.T) {
	setup(t)
	defer teardown()

        want := &Domain{
                Name: "example.com",
                TTL: 1800,
                ZoneFile: "Example zone file text...",
        }

        mux.HandleFunc("/domains", testServer(200, listDomainsExample))

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
                Name: "example.com",
                TTL: 1800,
                ZoneFile: "Example zone file text...",
        }

        u := fmt.Sprintf("/domains/%s", want.Name)
        mux.HandleFunc(u, testServer(200, domainExample))

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
                Name: "example.com",
                TTL: 1800,
                ZoneFile: "Example zone file text...",
        }

	mux.HandleFunc("/domains", testServer(202, domainExample))

        domain, err := client.CreateDomain(want.Name, "127.0.0.20")
        assertEqual(t, err, nil)
        assertEqual(t, domain.Name, want.Name)
        assertEqual(t, domain.TTL, want.TTL)
        assertEqual(t, domain.ZoneFile, want.ZoneFile)
}

func Test_DeleteDomain(t *testing.T) {
	setup(t)
	defer teardown()

        name := "example.com"
        u := fmt.Sprintf("/domains/%s", name)
        mux.HandleFunc(u, testServer(204, ""))

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
