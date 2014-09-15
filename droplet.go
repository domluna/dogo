package digitalocean

import (
	"fmt"
	"time"
)

const (
	DropletEndpoint = "droplets"
)

// Droplet respresents a DigitalOcean droplet.
type Droplet struct {
	ID          int       `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Memory      int       `json:"memory,omitempty"`
	VCPUS       int       `json:"vcpus,omitempty"`
	Disk        int       `json:"disk,omitempty"`
	Region      Region    `json:"region,omitempty"`
	Image       Image     `json:"image,omitempty"`
	Kernel      Kernel    `json:"kernel,omitempty"`
	Size        Size      `json:"size,omitempty"`
	Locked      bool      `json:"locked,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	Status      string    `json:"status,omitempty"`
	Networks    Networks  `json:"networks,omitempty"`
	BackupIDs   []int     `json:"backups_ids,omitempty"`
	SnapshotIDs []int     `json:"snapshot_ids,omitempty"`
	ActionIDs   []int     `json:"action_ids,omitempty"`
	Features    []string  `json:"features,omitempty"`
}

// IPV4 returns the ipv4 address of the droplet.
func (d *Droplet) IPV4() string {
	if len(d.Networks.V4) > 0 {
		return d.Networks.V4[0].IP
	}
	return ""
}

// IPV4 returns the ipv6 address of the droplet.
func (d *Droplet) IPV6() string {
	if len(d.Networks.V6) > 0 {
		return d.Networks.V6[0].IP
	}
	return ""
}

// SizeSlug returns the size of the droplet, ex: "512mb".
func (d *Droplet) SizeSlug() string {
	return d.Size.Slug
}

// ImageSlug return the name of the droplet's image, ex: "Ubuntu 13.10 x64 ... "
func (d *Droplet) ImageName() string {
	return d.Image.Name
}

// ImageID return the id of the droplet's image, ex: 3668014
func (d *Droplet) ImageID() int {
	return d.Image.ID
}

type Droplets []Droplet

// GetDroplets returns all users droplets, active or otherwise.
func (c *Client) ListDroplets() (Droplets, error) {
	s := struct {
		Droplets `json:"droplets,omitempty"`
	}{}
	err := c.get(DropletEndpoint, &s)
	if err != nil {
		return nil, err
	}
	return s.Droplets, nil
}

// GetDroplet return an individual droplet based on the passed id.
func (c *Client) GetDroplet(id int) (*Droplet, error) {
	u := fmt.Sprintf("%s/%d", DropletEndpoint, id)
	s := struct {
		Droplet `json:"droplet,omitempty"`
	}{}

	err := c.get(u, &s)
	if err != nil {
		return nil, err
	}
	return &s.Droplet, nil
}

// CreateDroplet creates a droplet based on based specs.
func (c *Client) CreateDroplet(params Params) (*Droplet, error) {
	s := struct {
		Droplet `json:"droplet,omitempty"`
	}{}
	err := c.post(DropletEndpoint, params, &s)
	if err != nil {
		return nil, err
	}
	return &s.Droplet, nil
}

// DestroyDroplet destroys a droplet. CAUTION - this is irreversible.
// There may be more appropriate options.
func (c *Client) DestroyDroplet(id int) error {
	u := fmt.Sprintf("%s/%d", DropletEndpoint, id)
	err := c.delete(u)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) ResizeDroplet(id int, size string) error {
	return c.DoAction(DropletEndpoint, id, Params{
		"type": "resize",
		"size": size,
	})
}

func (c *Client) RenameDroplet(id int, name string) error {
	return c.DoAction(DropletEndpoint, id, Params{
		"type": "resize",
		"name": name,
	})
}

func (c *Client) EnableIPV6(id int, size string) error {
	return c.DoAction(DropletEndpoint, id, Params{
		"type": "enable_ipv6",
	})

}

func (c *Client) EnablePrivateNetworking(id int) error {
	return c.DoAction(DropletEndpoint, id, Params{
		"type": "enable_private_networking",
	})
}

func (c *Client) PowerOffDroplet(id int) error {
	return c.DoAction(DropletEndpoint, id, Params{
		"type": "power_off",
	})
}

func (c *Client) PowerOnDroplet(id int) error {
	return c.DoAction(DropletEndpoint, id, Params{
		"type": "power_on",
	})
}
