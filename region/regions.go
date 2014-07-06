package region

import (
	"github.com/domluna/dogo/digitalocean"
)

const (
	Endpoint = digitalocean.BaseURL + "/regions"
)

// Region represents a DigitalOcean region.
type Region struct {
	Name      string   `json:"name,omitempty"`
	Slug      string   `json:"slug,omitempty"`
	Sizes     []string `json:"sizes,omitempty"`
	Available bool     `json:"available,omitempty"`
	Features  []string `json:"features,omitempty"`
}

type Regions []Region

type Client struct {
	client digitalocean.Client
}

func NewClient(token string) *Client {
	return &Client{digitalocean.NewClient(token)}
}

// GetRegions gets all current available regions a droplet may be created in.
func (c *Client) GetAll() (Regions, error) {
	s := struct {
		Regions           `json:"regions,omitempty"`
		digitalocean.Meta `json:"meta,omitempty"`
	}{}
	err := c.client.Get(Endpoint, &s)
	if err != nil {
		return s.Regions, err
	}
	return s.Regions, nil
}
