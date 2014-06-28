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
	var d Domains
	err := dc.Client.Get(DomainsEndpoint, d)
	if err != nil {
		return d, err
	}
	return d, nil
}

func (dc *DomainClient) GetDomain(name string) (Domain, error) {
	u := fmt.Sprintf("%s/%s", DomainsEndpoint, name)
	var d Domain
	err := dc.Client.Get(u, d)
	if err != nil {
		return d, err
	}
	return d, nil
}

func (dc *DomainClient) CreateDomain(name string, ip string) (Domain, error) {
	payload := map[string]interface{}{
		"name":       name,
		"ip_address": ip,
	}

	var d Domain
	err := dc.Client.Post(DomainsEndpoint, payload, d)
	if err != nil {
		return d, err
	}
	return d, nil
}

func (dc *DomainClient) DelDomain(name string) error {
	u := fmt.Sprintf("%s/%s", DomainsEndpoint, name)
	err := dc.Client.Del(u)
	if err != nil {
		return err
	}
	return nil
}
