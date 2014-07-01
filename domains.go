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

type DomainClient struct {
	client Client
}

func (dc *DomainClient) GetAll() (Domains, error) {
	s := struct {
		Domains `json:"domains,omitempty"`
		Meta    `json:"meta,omitempty"`
	}{}
	err := dc.client.Get(DomainsEndpoint, &s)
	if err != nil {
		return nil, err
	}
	return s.Domains, nil
}

func (dc *DomainClient) Get(name string) (Domain, error) {
	u := fmt.Sprintf("%s/%s", DomainsEndpoint, name)
	s := struct {
		Domain `json:"domains,omitempty"`
	}{}
	err := dc.client.Get(u, &s)
	if err != nil {
		return s.Domain, err
	}
	return s.Domain, nil
}

func (dc *DomainClient) Create(name string, ip string) (Domain, error) {
	s := struct {
		Domain `json:"domains,omitempty"`
	}{}
	payload := map[string]interface{}{
		"name":       name,
		"ip_address": ip,
	}
	err := dc.client.Post(DomainsEndpoint, payload, &s)
	if err != nil {
		return s.Domain, err
	}
	return s.Domain, nil
}

func (dc *DomainClient) Delete(name string) error {
	u := fmt.Sprintf("%s/%s", DomainsEndpoint, name)
	err := dc.client.Del(u)
	if err != nil {
		return err
	}
	return nil
}
