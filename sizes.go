package dogo

// Representation for the size of a DigitalOcean droplet.
type Size struct {
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	Slug         string   `json:"slug"`
	Memory       int      `json:"memory"`
	Vcpus        int      `json:"vcpus"`
	Disk         int      `json:"disk"`
	Transfer     int      `json:"transfer"`
	PriceHourly  string   `json:"price_hourly"`
	PriceMonthly string   `json:"price_monthly"`
	Regions      []string `json:"regions"`
}

// SizesMap is a mapping between the slug
// representation of a droplet size to it's
// id.
var SizesMap = map[string]int{
	"512MB": 66,
	"1GB":   63,
	"2GB":   62,
	"4GB":   64,
	"8GB":   65,
	"16GB":  61,
	"32GB":  60,
	"48GB":  70,
	"64GB":  69,
}

// GetSizes returns all currently available droplet sizes.
func (c *Client) GetSizes() ([]Size, error) {
	resp, err := c.send(SizesEndpoint, nil, nil)
	if err != nil {
		return nil, err
	}
	return resp.Sizes, nil
}
