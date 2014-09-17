package dogo

const RegionEndpoint = "regions"

// Region represents a DigitalOcean region.
type Region struct {
	Slug      string   `json:"slug,omitempty"`
	Name      string   `json:"name,omitempty"`
	Sizes     []string `json:"sizes,omitempty"`
	Available bool     `json:"available,omitempty"`
	Features  []string `json:"features,omitempty"`
}

// Regions is a list of type Region.
type Regions []*Region

// GetRegions gets all current available regions a droplet may be created in.
func (c *Client) ListRegions() (Regions, error) {
	s := struct {
		Regions `json:"regions,omitempty"`
	}{}
	err := c.get(RegionEndpoint, &s)
	if err != nil {
		return nil, err
	}
	return s.Regions, nil
}
