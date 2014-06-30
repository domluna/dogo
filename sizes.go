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

type SizeClient struct {
	Client
}

// GetSizes returns all currently available droplet sizes.
func (sc *SizeClient) GetSizes() (Sizes, error) {
	var s Sizes
	req, err := sc.Client.Get(SizesEndpoint)
	if err != nil {
		return s, err
	}

	sc.Client.DoRequest(req, &s)
	if err != nil {
		return s, err
	}
	return s, nil
}
