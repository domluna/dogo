package dogo

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

// GetSizes returns all currently available droplet sizes.
func (c *Client) GetSizes() (Sizes, error) {
	var s Sizes
	err := c.Get(SizesEndpoint, s)
	if err != nil {
		return s, err
	}
	return s, nil
}
