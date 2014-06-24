package dogo

// Representation for a DigitalOcean Image.
type Image struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Distribution string `json:"distribution"`
	Slug         string `json:"slug"`
	Public       bool   `json:"public"`
}

// GetImages returns DigitalOcean images, filter can either be
// "my_images" or "global".
//
// If filter is set to "my_images" user snapshots will be returned.
//
// If filter is set to "global" all default images will be returned.
func (c *Client) GetMyImages() ([]Image, error) {
	resp, err := c.send(ImagesEndpoint, nil, Params{
		"filter": "my_images",
	})
	if err != nil {
		return resp.Images, err
	}
	return resp.Images, nil
}

func (c *Client) GetAllImages() ([]Image, error) {
	resp, err := c.send(ImagesEndpoint, nil, Params{
		"filter": "global",
	})
	if err != nil {
		return resp.Images, err
	}
	return resp.Images, nil
}
