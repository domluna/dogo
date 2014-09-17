package dogo

// Size is a representation for the size of a DigitalOcean droplet.
type Size struct {
	Slug         string   `json:"slug,omitempty"`
	Memory       int      `json:"memory,omitempty"`
	VCPUS        int      `json:"vcpus,omitempty"`
	Disk         int      `json:"disk,omitempty"`
	Transfer     int      `json:"transfer,omitempty"`
	PriceMonthly float32  `json:"price_monthly,omitempty"`
	PriceHourly  float32  `json:"price_hourly,omitempty"`
	Regions      []string `json:"regions,omitempty"`
}

type Sizes []*Size

// ListSizes returns all currently available droplet sizes.
func (c *Client) ListSizes() (Sizes, error) {
	s := struct {
		Sizes `json:"sizes,omitempty"`
	}{}
	err := c.get(SizeEndpoint, &s)
	if err != nil {
		return nil, err
	}
	return s.Sizes, nil
}
