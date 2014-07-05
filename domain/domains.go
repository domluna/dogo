package dogo

import (
	"fmt"
)

type Domain struct {
	Name     string `json:"name,omitempty"`
	TTL      int    `json:"ttl,omitempty"`
	ZoneFile string `json:"zone_file,omitempty"`
}

type Domains []Domain

type Client struct {
	client Client
}

func (c *Client) GetAll() (Domains, error) {
	s := struct {
		Domains `json:"domains,omitempty"`
		Meta    `json:"meta,omitempty"`
	}{}
	err := c.client.Get(DomainsEndpoint, &s)
	if err != nil {
		return nil, err
	}
	return s.Domains, nil
}

func (c *Client) Get(name string) (Domain, error) {
	u := fmt.Sprintf("%s/%s", DomainsEndpoint, name)
	s := struct {
		Domain `json:"domains,omitempty"`
	}{}
	err := c.client.Get(u, &s)
	if err != nil {
		return s.Domain, err
	}
	return s.Domain, nil
}

func (c *Client) Create(name string, ip string) (Domain, error) {
	s := struct {
		Domain `json:"domains,omitempty"`
	}{}
	payload := map[string]interface{}{
		"name":       name,
		"ip_address": ip,
	}
	err := c.client.Post(DomainsEndpoint, payload, &s)
	if err != nil {
		return s.Domain, err
	}
	return s.Domain, nil
}

func (c *Client) Delete(name string) error {
	u := fmt.Sprintf("%s/%s", DomainsEndpoint, name)
	err := c.client.Del(u)
	if err != nil {
		return err
	}
	return nil
}
