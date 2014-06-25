package dogo

// Region represents a DigitalOcean region.
type Region struct {
	Name      string   `json:"name"`
	Slug      string   `json:"slug"`
	Sizes     []string `json:"sizes"`
	Available bool     `json:"available"`
}

// RegionsMap is a mapping between the slug
// representation of the region and it's id.
//
// Note that some regions listed may actually not be
// currently available.
var RegionsMap = map[string]int{
	"nyc1": 1,
	"ams1": 2,
	"sfo1": 3,
	"nyc2": 4,
	"ams2": 5,
	"sgp1": 6,
}

// GetRegions gets all current available regions a droplet may be created in.
func (c *Client) GetRegions() ([]Region, error) {
	resp, err := c.send(RegionsEndpoint, nil, nil)
	if err != nil {
		return resp.Regions, err
	}
	return resp.Regions, nil
}
