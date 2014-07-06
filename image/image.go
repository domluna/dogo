package image

import (
	"fmt"
	"time"

	"github.com/domluna/dogo/digitalocean"
)

const (
	Endpoint = digitalocean.BaseURL + "/images"
)

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

type Images []Image

type Client struct {
	client digitalocean.Client
}

func NewClient(token string) *Client {
	return &Client{digitalocean.NewClient(token)}
}

// GetMyImages gets all custom images/snapshots.
func (c *Client) GetAll() (Images, error) {
	s := struct {
		Images            `json:"images,omitempty"`
		digitalocean.Meta `json:"meta,omitempty"`
	}{}
	err := c.client.Get(Endpoint, &s)
	if err != nil {
		return s.Images, err
	}
	return s.Images, nil
}

// GetMyImages gets all custom images/snapshots.
func (c *Client) Get(v interface{}) (Image, error) {
	u := fmt.Sprintf("%s/%v", Endpoint, v)
	s := struct {
		Image             `json:"images,omitempty"`
		digitalocean.Meta `json:"meta,omitempty"`
	}{}
	err := c.client.Get(u, &s)
	if err != nil {
		return s.Image, err
	}
	return s.Image, nil
}

// GetMyImages gets all custom images/snapshots.
func (c *Client) Delete(id int) error {
	u := fmt.Sprintf("%s/%d", Endpoint, id)
	err := c.client.Del(u)
	if err != nil {
		return err
	}
	return nil
}

// GetMyImages gets all custom images/snapshots.
func (c *Client) Update(id int, name string) (Image, error) {
	u := fmt.Sprintf("%s/%d", Endpoint, id)
	s := struct {
		Image `json:"image,omitempty"`
	}{}
	payload := map[string]interface{}{
		"name": name,
	}
	err := c.client.Put(u, payload, &s)
	if err != nil {
		return s.Image, err
	}
	return s.Image, nil
}

func (c *Client) DoAction(id int, params map[string]interface{}) error {
	u := fmt.Sprintf("%s/%d", Endpoint, id)
	err := c.client.Post(u, params, nil)
	if err != nil {
		return err
	}
	return nil
}
