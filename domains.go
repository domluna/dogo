package dogo

import (
	"fmt"
)

type Domain struct {
	Name     string `json:"name,omitempty"`
	TTL      int    `json:"ttl,omitempty"`
	ZoneFile string `json"zone_file,omitempty"`
}

type Domains []Domain

type DomainClient struct {
	Client
}

func (dc *DomainClient) GetDomains() (Domains, error) {
	s := struct{
		Domains `json:"domains,omitempty"`
		Meta `json:"meta,omitempty"`
	}{}
	err := dc.Get(DomainsEndpoint, &s)
	if err != nil {
		return nil, err
	}
	return s.Domains, nil
}

func (dc *DomainClient) GetDomain(name string) (Domain, error) {
	u := fmt.Sprintf("%s/%s", DomainsEndpoint, name)
	s := struct{
		Domain `json:"domains,omitempty"`
	}{}
	err := dc.Get(u, &s)
	if err != nil {
		return s.Domain, err
	}
	return s.Domain, nil
}

func (dc *DomainClient) CreateDomain(name string, ip string) (Domain, error) {
	s := struct{
		Domain `json:"domains,omitempty"`
	}{}
	payload := map[string]interface{}{
		"name":       name,
		"ip_address": ip,
	}
	err := dc.Post(DomainsEndpoint, payload, &s)
	if err != nil {
		return s.Domain, err
	}
	return s.Domain, nil
}

func (dc *DomainClient) DelDomain(name string) error {
	u := fmt.Sprintf("%s/%s", DomainsEndpoint, name)
	err := dc.Del(u)
	if err != nil {
		return err
	}
	return nil
}
