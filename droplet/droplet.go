package droplet

import (
	"fmt"

	"github.com/domluna/dogo/digitalocean"
	"github.com/domluna/dogo/image"
	"github.com/domluna/dogo/region"
	"github.com/domluna/dogo/size"
)

const (
	Endpoint = digitalocean.BaseURL + "/droplets"
)

// Droplet respresents a DigitalOcean droplet.
type Droplet struct {
	ID          int           `json:"id,omitempty"`
	Name        string        `json:"name,omitempty"`
	Region      region.Region `json:"region,omitempty"`
	Image       image.Image   `json:"image,omitempty"`
	Kernel      image.Kernel  `json:"kernel,omitempty"`
	Size        size.Size     `json:"size,omitempty"`
	Locked      bool          `json:"locked,omitempty"`
	Status      string        `json:"status,omitempty"`
	Networks    Networks      `json:"networks,omitempty"`
	BackupIDs   []int         `json:"backups_ids,omitempty"`
	SnapshotIDs []int         `json:"snapshot_ids,omitempty"`
	ActionIDs   []int         `json:"action_ids,omitempty"`
}

type Droplets []Droplet

type Client struct {
	client digitalocean.Client
}

func NewClient(token string) *Client {
	return &Client{digitalocean.NewClient(token)}
}

// GetDroplets returns all users droplets, active or otherwise.
func (c *Client) GetAll() (Droplets, error) {
	s := struct {
		Droplets          `json:"droplets,omitempty"`
		digitalocean.Meta `json:"meta,omitempty"`
	}{}
	err := c.client.Get(Endpoint, &s)
	if err != nil {
		return nil, err
	}
	return s.Droplets, nil
}

// GetDroplet return an individual droplet based on the passed id.
func (c *Client) Get(id int) (Droplet, error) {
	u := fmt.Sprintf("%s/%d", Endpoint, id)
	s := struct {
		Droplet `json:"droplet,omitempty"`
	}{}

	err := c.client.Get(u, &s)
	if err != nil {
		return s.Droplet, err
	}
	return s.Droplet, nil
}

// CreateDroplet creates a droplet based on based specs.
func (c *Client) Create(params digitalocean.Params) (Droplet, error) {
	s := struct {
		Droplet `json:"droplet,omitempty"`
	}{}
	err := c.client.Post(Endpoint, params, &s)
	if err != nil {
		return s.Droplet, err
	}
	return s.Droplet, nil
}

// DestroyDroplet destroys a droplet. CAUTION - this is irreversible.
// There may be more appropriate options.
func (c *Client) Destroy(id int) error {
	u := fmt.Sprintf("%s/%d", Endpoint, id)
	err := c.client.Delete(u)
	if err != nil {
		return err
	}
	return nil
}

// DoAction performs an action on a droplet with the passed id, type of action and its
// required params are described in the DigitalOcean API.
//	https://developers.digitalocean.com/v2/#droplet-actions
//
// An example of some params:
//	params := digitalocean.Params{
//		"type": "resize",
//		"size": "1024mb",
//	}
//
// The above example specifies the type of action, in this case resizing and
// the additional param in this case the size to resize to "1024mb".
//
// Params will sometimes only require the type of action and no additional params.
//
//
func (c *Client) DoAction(id int, params digitalocean.Params) error {
	u := fmt.Sprintf("%s/%d/actions", Endpoint, id)
	err := c.client.Post(u, params, nil)
	if err != nil {
		return err
	}
	return nil
}
