package dogo

// Region represents a DigitalOcean region.
type Region struct {
	Name      string   `json:"name,omitempty"`
	Slug      string   `json:"slug,omitempty"`
	Sizes     []string `json:"sizes,omitempty"`
	Available bool     `json:"available,omitempty"`
	Features  []string `json"features,omitempty"`
}

type Regions []Region

type RegionClient struct {
	Client
}

// GetRegions gets all current available regions a droplet may be created in.
func (rc *RegionClient) GetRegions() (Regions, error) {
	s := struct {
		Regions `json"regions,omitempty"`
		Meta    `json"meta,omitempty"`
	}{}
	err := rc.Get(RegionsEndpoint, &s)
	if err != nil {
		return s.Regions, err
	}
	return s.Regions, nil
}
