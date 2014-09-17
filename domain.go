package dogo

import (
	"fmt"
)

const (
	DomainEndpoint = "domains"
)

// Domain is a representation of a DigitalOcean domain.
type Domain struct {
	// Name of the Domain.
	Name string `json:"name,omitempty"`

	// Defines the time frame that clients can
	// cache queried information before a refresh should
	// be requests.
	TTL int `json:"ttl,omitempty"`

	// Contains the complete contents of the
	// zone file.
	ZoneFile string `json:"zone_file,omitempty"`
}

// Domains is a list of Domain.
type Domains []*Domain

type CreateDomainOpts struct {
	Name      string `json:"name"`
	IPAddress string `json:"ip_address"`
}

func (c *Client) ListDomains() (Domains, error) {
	s := struct {
		Domains `json:"domains,omitempty"`
	}{}
	err := c.get(DomainEndpoint, &s)
	if err != nil {
		return nil, err
	}
	return s.Domains, nil
}

func (c *Client) GetDomain(name string) (*Domain, error) {
	u := fmt.Sprintf("%s/%s", DomainEndpoint, name)
	s := struct {
		Domain `json:"domain,omitempty"`
	}{}
	err := c.get(u, &s)
	if err != nil {
		return nil, err
	}
	return &s.Domain, nil
}

func (c *Client) CreateDomain(opts *CreateDomainOpts) (*Domain, error) {
	s := struct {
		Domain `json:"domain,omitempty"`
	}{}
	err := c.post(DomainEndpoint, opts, &s)
	if err != nil {
		return nil, err
	}
	return &s.Domain, nil
}

func (c *Client) DeleteDomain(name string) error {
	u := fmt.Sprintf("%s/%s", DomainEndpoint, name)
	err := c.delete(u)
	if err != nil {
		return err
	}
	return nil
}
