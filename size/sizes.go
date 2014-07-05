package size

// Representation for the size of a DigitalOcean droplet.
type Size struct {
	ID           int      `json:"id,omitempty"`
	Name         string   `json:"name,omitempty"`
	Slug         string   `json:"slug,omitempty"`
	Memory       int      `json:"memory,omitempty"`
	Vcpus        int      `json:"vcpus,omitempty"`
	Disk         int      `json:"disk,omitempty"`
	Transfer     int      `json:"transfer,omitempty"`
	PriceHourly  string   `json:"price_hourly,omitempty"`
	PriceMonthly string   `json:"price_monthly,omitempty"`
	Regions      []string `json:"regions,omitempty"`
}

type Sizes []Size

type Client struct {
	client Client
}

// GetSizes returns all currently available droplet sizes.
func (c *Client) GetAll() (Sizes, error) {
	s := struct {
		Sizes `json:"sizes,omitempty"`
		Meta  `json:"meta,omitempty"`
	}{}
	err := c.client.Get(SizesEndpoint, &s)
	if err != nil {
		return s.Sizes, err
	}
	return s.Sizes, nil
}
