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
	req, err := dc.Client.Get(DomainsEndpoint)
	if err != nil {
		return d, err
	}

	err = dc.Client.DoRequest(req, &d)
	if err != nil {
		return d, err
	}
	return d, nil
}

func (dc *DomainClient) GetDomain(name string) (Domain, error) {
	u := fmt.Sprintf("%s/%s", DomainsEndpoint, name)
	var d Domain
	req, err := dc.Client.Get(u)
	if err != nil {
		return d, err
	}

	err = dc.Client.DoRequest(req, &d)
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
	req, err := dc.Client.Post(DomainsEndpoint, payload)
	if err != nil {
		return d, err
	}

	err = dc.Client.DoRequest(req, &d)
	if err != nil {
		return d, err
	}
	return d, nil
}

func (dc *DomainClient) DelDomain(name string) error {
	u := fmt.Sprintf("%s/%s", DomainsEndpoint, name)
	req, err := dc.Client.Del(u)
	if err != nil {
		return err
	}

	err = dc.Client.DoRequest(req, nil)
	if err != nil {
		return err
	}
	return nil
}
