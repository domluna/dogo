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
	err := c.client.Delete(u)
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
	payload := digitalocean.Params{
		"name": name,
	}
	err := c.client.Put(u, payload, &s)
	if err != nil {
		return s.Image, err
	}
	return s.Image, nil
}

// DoAction performs an action on a droplet with the passed id, type of action and its
// required params are described in the DigitalOcean API.
//	https://developers.digitalocean.com/v2/#image-actions
//
// An example of some params:
//	params := digitalocean.Params{
//		"type": "transfer",
//		"region": "nyc2",
//	}
//
// The above example specifies the type of action, in this case resizing and
// the additional param in this case the size to resize to "1024mb".
//
// Params will sometimes only require the type of action and no additional params.
//
//
func (c *Client) DoAction(id int, params digitalocean.Params) error {
	u := fmt.Sprintf("%s/%d", Endpoint, id)
	err := c.client.Post(u, params, nil)
	if err != nil {
		return err
	}
	return nil
}
