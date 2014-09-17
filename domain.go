package dogo

import (
	"fmt"
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

type Domains []*Domain

// CreateDomainOpts contains options used when creating a new domain.
type CreateDomainOpts struct {
	// Name of the domain.
	Name string `json:"name"`

	// Address domain will point to.
	IPAddress string `json:"ip_address"`
}

// ListDomains retrieves all user domains.
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

// GetDomain retrieves the domain given its name.
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

// CreateDomain creates a domain name.
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

// DeleteDomain deletes the passed domain name.
func (c *Client) DeleteDomain(name string) error {
	u := fmt.Sprintf("%s/%s", DomainEndpoint, name)
	err := c.delete(u)
	if err != nil {
		return err
	}
	return nil
}
