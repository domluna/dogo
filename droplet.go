package dogo

import (
	"fmt"
)

// Networks represents a Droplet Network.
// This is represented as two slices of
// ipv4 and ipv6 addresses.
type Networks struct {
	V4 []*V4 `json:"v4,omitempty"`
	V6 []*V6 `json:"v6,omitempty"`
}

// V4 represents a ipv4 address.
type V4 struct {
	IPAddress string `json:"ip_address,omitempty"`
	Netmask   string `json:"netmask,omitempty"`
	Gateway   string `json:"gateway,omitempty"`
	Type      string `json:"type,omitempty"`
}

// V6 represents a ipv6 address.
type V6 struct {
	IPAddress string `json:"ip_address,omitempty"`
	Cidr      int    `json:"cidr,omitempty"`
	Gateway   string `json:"gateway,omitempty"`
	Type      string `json:"type,omitempty"`
}

// Droplet respresents a DigitalOcean droplet.
type Droplet struct {
	ID          int       `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Memory      int       `json:"memory,omitempty"`
	VCPUS       int       `json:"vcpus,omitempty"`
	Disk        int       `json:"disk,omitempty"`
	Region      *Region   `json:"region,omitempty"`
	Image       *Image    `json:"image,omitempty"`
	Kernel      *Kernel   `json:"kernel,omitempty"`
	Size        *Size     `json:"size,omitempty"`
	Locked      bool      `json:"locked,omitempty"`
	CreatedAt   string    `json:"created_at,omitempty"`
	Status      string    `json:"status,omitempty"`
	Networks    *Networks `json:"networks,omitempty"`
	BackupIDs   []int     `json:"backup_ids,omitempty"`
	SnapshotIDs []int     `json:"snapshot_ids,omitempty"`
	ActionIDs   []int     `json:"action_ids,omitempty"`
	Features    []string  `json:"features,omitempty"`
}

// IPV4Addr returns the ipv4 address of the droplet.
func (d *Droplet) IPV4Addr() string {
	if len(d.Networks.V4) > 0 {
		for _, net := range d.Networks.V4 {
			if net.Type == "public" {
				return net.IPAddress
			}
		}
	}
	return ""
}

// IPV6Addr returns the ipv6 address of the droplet.
func (d *Droplet) IPV6Addr() string {
	if len(d.Networks.V6) > 0 {
		for _, net := range d.Networks.V6 {
			if net.Type == "public" {
				return net.IPAddress
			}
		}
	}
	return ""
}

// SizeSlug returns the size of the droplet,
// ex: "512mb".
func (d *Droplet) SizeSlug() string {
	return d.Size.Slug
}

// ImageName return the name of the droplet's image,
// ex: "Ubuntu 13.10 x64 ... ".
func (d *Droplet) ImageName() string {
	return d.Image.Name
}

// ImageID return the id of the droplet's image,
// ex: 3668014.
func (d *Droplet) ImageID() int {
	return d.Image.ID
}

// KernelName returns the name of the kernel,
// ex: "Ubuntu 14.04 x32 vmlinuz-3.13.0-24-generic".
func (d *Droplet) KernelName() string {
	return d.Kernel.Name
}

// RegionSlug returns the region slug,
// ex: "nyc3" for New York 3.
func (d *Droplet) RegionSlug() string {
	return d.Region.Slug
}

// Droplets is a list of Droplet.
type Droplets []*Droplet

// CreateDropletOpts is a utility object used when
// creating a droplet.
//
// Name, Region, Size and Image are required.
type CreateDropletOpts struct {
	Name   string `json:"name"`
	Region string `json:"region"`
	Size   string `json:"size"`

	// Image can either be an id or a slug.
	Image int `json:"image"`

	// The key can be either an id or a fingerprint
	Keys []string `json:"ssh_keys"`

	Backups           bool   `json:"backups"`
	IPV6              bool   `json:"ipv6"`
	PrivateNetworking bool   `json:"private_networking"`
	UserData          string `json:"user_data"`
}

// ListDroplets returns all users droplets, active or otherwise.
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
func (c *Client) CreateDroplet(opts *CreateDropletOpts) (*Droplet, error) {
	s := struct {
		Droplet `json:"droplet,omitempty"`
	}{}
	err := c.post(DropletEndpoint, opts, &s)
	if err != nil {
		return nil, err
	}
	return &s.Droplet, nil
}

// DeleteDroplet destroys a droplet. CAUTION - this is irreversible.
// There may be more appropriate options.
func (c *Client) DeleteDroplet(id int) error {
	u := fmt.Sprintf("%s/%d", DropletEndpoint, id)
	err := c.delete(u)
	if err != nil {
		return err
	}
	return nil
}

// ListKernels retrieves all kernels for a particular Droplet.
func (c *Client) ListKernels(id int) (Kernels, error) {
	u := fmt.Sprintf("%s/%d/kernels", DropletEndpoint, id)
	s := struct {
		Kernels `json:"kernels,omitempty"`
	}{}
	err := c.get(u, &s)
	if err != nil {
		return nil, err
	}
	return s.Kernels, nil
}

// ListSnapshots retrieves all snapshots for a particular Droplet.
func (c *Client) ListSnapshots(id int) (Snapshots, error) {
	u := fmt.Sprintf("%s/%d/snapshots", DropletEndpoint, id)
	s := struct {
		Snapshots `json:"snapshots,omitempty"`
	}{}
	err := c.get(u, &s)
	if err != nil {
		return nil, err
	}
	return s.Snapshots, nil
}

// ListBackups retrieves all backups for a particular Droplet.
func (c *Client) ListBackups(id int) (Backups, error) {
	u := fmt.Sprintf("%s/%d/backups", DropletEndpoint, id)
	s := struct {
		Backups `json:"backups,omitempty"`
	}{}
	err := c.get(u, &s)
	if err != nil {
		return nil, err
	}
	return s.Backups, nil
}
