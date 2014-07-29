package size

import "github.com/domluna/dogo/digitalocean"

const (
	Endpoint = digitalocean.BaseURL + "/sizes"
)

// Representation for the size of a DigitalOcean droplet.
type Size struct {
	ID           int      `json:"id,omitempty"`
	Name         string   `json:"name,omitempty"`
	Slug         string   `json:"slug,omitempty"`
	Memory       int      `json:"memory,omitempty"`
	Vcpus        int      `json:"vcpus,omitempty"`
	Disk         int      `json:"disk,omitempty"`
	Transfer     int      `json:"transfer,omitempty"`
	PriceHourly  float32  `json:"price_hourly,omitempty"`
	PriceMonthly float32  `json:"price_monthly,omitempty"`
	Regions      []string `json:"regions,omitempty"`
}

type Sizes []Size

type Client struct {
	client digitalocean.Client
}

func NewClient(token string) *Client {
	return &Client{digitalocean.NewClient(token)}
}

// GetSizes returns all currently available droplet sizes.
func (c *Client) GetAll() (Sizes, error) {
	s := struct {
		Sizes             `json:"sizes,omitempty"`
	}{}
	err := c.client.Get(Endpoint, &s)
	if err != nil {
		return s.Sizes, err
	}
	return s.Sizes, nil
}
