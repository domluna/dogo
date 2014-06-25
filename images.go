package dogo

// Kernel is a DigitalOcean Kernel.
type Kernel struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Version string `json:"version"`
}

// Snapshot is a DigitalOcean Snapshot/Backup.
type Snapshot struct {
	ID        int      `json:"id"`
	Name      string   `json:"name"`
	Dist      string   `json:"distribution"`
	Slug      string   `json:"slug"`
	Public    bool     `json:"public"`
	Regions   []string `json:"regions"`
	ActionIDs []int    `json:"action_ids"`
}

// Representation for a DigitalOcean Image.
type Image struct {
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	Distribution string   `json:"distribution"`
	Slug         string   `json:"slug"`
	Public       bool     `json:"public"`
	Regions      []string `json"regions"`
}

// GetMyImages gets all custom images/snapshots.
func (c *Client) GetMyImages() ([]Image, error) {
	resp, err := c.send(ImagesEndpoint, nil, Params{
		"filter": "my_images",
	})
	if err != nil {
		return resp.Images, err
	}
	return resp.Images, nil
}

// GetAllImages gets all default images.
func (c *Client) GetAllImages() ([]Image, error) {
	resp, err := c.send(ImagesEndpoint, nil, Params{
		"filter": "global",
	})
	if err != nil {
		return resp.Images, err
	}
	return resp.Images, nil
}
