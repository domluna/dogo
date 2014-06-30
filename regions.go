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
	var r Regions
	req, err := rc.Client.Get(RegionsEndpoint)
	if err != nil {
		return r, err
	}

	err = rc.Client.DoRequest(req, &r)
	if err != nil {
		return r, err
	}
	return r, nil
}
