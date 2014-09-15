package digitalocean

import (
	"fmt"
	"time"
)

const ImageEndpoint = "images"

// Kernel is a DigitalOcean Kernel.
type Kernel struct {
	ID      int    `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Version string `json:"version,omitempty"`
}

// Snapshot is a DigitalOcean Snapshot/Backup.
type Snapshot struct {
	ID        int      `json:"id,omitempty"`
	Name      string   `json:"name,omitempty"`
	Dist      string   `json:"distribution,omitempty"`
	Slug      string   `json:"slug,omitempty,omitempty"`
	Public    bool     `json:"public,omitempty"`
	Regions   []string `json:"regions,omitempty"`
	ActionIDs []int    `json:"action_ids,omitempty"`
}

// Representation for a DigitalOcean Image.
type Image struct {
	ID           int       `json:"id,omitempty"`
	Name         string    `json:"name,omitempty"`
	Distribution string    `json:"distribution,omitempty"`
	Slug         string    `json:"slug,omitempty,omitempty"`
	Public       bool      `json:"public,omitempty"`
	Regions      []string  `json:"regions,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
}

// Images is a list of type Image.
type Images []*Image

// GetMyImages gets all custom images/snapshots.
func (c *Client) ListImages() (Images, error) {
	s := struct {
		Images `json:"images,omitempty"`
	}{}
	err := c.get(ImageEndpoint, &s)
	if err != nil {
		return nil, err
	}
	return s.Images, nil
}

// GetMyImages gets all custom images/snapshots.
func (c *Client) GetImage(v interface{}) (Image, error) {
	u := fmt.Sprintf("%s/%v", ImageEndpoint, v)
	s := struct {
		Image `json:"images,omitempty"`
	}{}
	err := c.get(u, &s)
	if err != nil {
		return s.Image, err
	}
	return s.Image, nil
}

// GetMyImages gets all custom images/snapshots.
func (c *Client) Delete(id int) error {
	u := fmt.Sprintf("%s/%d", ImageEndpoint, id)
	err := c.delete(u)
	if err != nil {
		return err
	}
	return nil
}

// GetMyImages gets all custom images/snapshots.
func (c *Client) Update(id int, name string) (Image, error) {
	u := fmt.Sprintf("%s/%d", ImageEndpoint, id)
	s := struct {
		Image `json:"image,omitempty"`
	}{}
	payload := Params{
		"name": name,
	}
	err := c.put(u, payload, &s)
	if err != nil {
		return s.Image, err
	}
	return s.Image, nil
}
