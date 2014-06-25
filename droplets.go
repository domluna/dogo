package dogo

import (
	"fmt"
	"strconv"
)

// Droplet respresents a digitalocean droplet.
// TODO: make networks and kernel types
type Droplet struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Region      `json:"region"`
	Image       `json:"image"`
	Kernel      string `json:"kernel"`
	Size        `json:"size"`
	Locked      bool   `json:"locked"`
	Status      string `json:"status"`
	Networks    string `json:"networks"`
	BackupIDs   []int  `json:"backups_ids"`
	SnapshotIDs []int  `json:"snapshot_ids"`
	ActionIDs   []int  `json:"action_ids"`
}

// GetDroplets returns all users droplets, active or otherwise.
func (c *Client) GetDroplets() ([]Droplet, error) {
	resp, err := c.send(DropletsEndpoint, nil, nil)
	if err != nil {
		return resp.Droplets, err
	}
	return resp.Droplets, nil

}

// GetDroplet return an individual droplet based on the passed id.
func (c *Client) GetDroplet(id int) (Droplet, error) {
	resp, err := c.send(DropletsEndpoint, id, nil)
	if err != nil {
		return *resp.Droplet, err
	}
	return *resp.Droplet, nil
}

// CreateDroplet creates a droplet based on based specs.
func (c *Client) CreateDroplet(d Droplet, keys []int, privateNet bool) (Droplet, error) {
	// Create a string of the key ids
	var keyStr string
	for _, k := range keys {
		ks := strconv.Itoa(k)
		keyStr += ks + ","
	}
	resp, err := c.send(DropletsEndpoint, "new", Params{
		"name":               d.Name,
		"size_id":            d.SizeID,
		"image_id":           d.ImageID,
		"region_id":          d.RegionID,
		"backups_enabled":    d.BackupsActive,
		"ssh_key_ids":        keyStr,
		"private_networking": privateNet,
	})
	if err != nil {
		return *resp.Droplet, err
	}
	return *resp.Droplet, nil
}

// DestroyDroplet destroys a droplet. CAUTION - this is irreversible.
// There may be more appropriate options.
func (c *Client) DestroyDroplet(id int) error {
	_, err := c.send(DropletsEndpoint, fmt.Sprintf("%d/destroy", id), nil)
	if err != nil {
		return err
	}
	return nil
}

// ResizeDroplet droplet resizes a droplet. Sizes are based on
// the digitalocean sizes api.
func (c *Client) ResizeDroplet(id, sizeID int) error {
	_, err := c.send(DropletsEndpoint, fmt.Sprintf("%d/resize", id), Params{
		"size_id": sizeID,
	})
	if err != nil {
		return err
	}
	return nil
}

// RebootDroplet reboots the a droplet. This is the preferred method
// to use if a server is not responding.
func (c *Client) RebootDroplet(id int) error {
	_, err := c.send(DropletsEndpoint, fmt.Sprintf("%d/reboot", id), nil)
	if err != nil {
		return err
	}
	return nil
}

// RebootDroplet rebuilds a droplet with a default image. This can be
// useful if you want to use a different image but keep the ip address
// of the droplet.
func (c *Client) RebuildDroplet(id, imageID int) error {
	_, err := c.send(DropletsEndpoint, fmt.Sprintf("%d/rebuild", id), Params{
		"image_id": imageID,
	})
	if err != nil {
		return err
	}
	return nil
}

// StopDroplet powers off a running droplet, the droplet will remain
// in your account.
func (c *Client) PowerOffDroplet(id int) error {
	_, err := c.send(DropletsEndpoint, fmt.Sprintf("%d/power_off", id), nil)
	if err != nil {
		return err
	}
	return nil
}

// StartDroplet powers on a powered off droplet.
func (c *Client) PowerOnDroplet(id int) error {
	_, err := c.send(DropletsEndpoint, fmt.Sprintf("%d/power_on", id), nil)
	if err != nil {
		return err
	}
	return nil
}

// SnapshotDroplet allows you to take a snapshot of a droplet once it is
// powered off. Be aware this may reboot the droplet.
func (c *Client) SnapshotDroplet(id int, name string) error {
	_, err := c.send(DropletsEndpoint, fmt.Sprintf("%d/snapshot", id), Params{
		"name": name,
	})
	if err != nil {
		return err
	}
	return nil
}

// RestoreDroplet allows you to restore a droplet to a previous image
// or snapshot. This will be a mirror copy of the image or snapshot to
// your droplet.
func (c *Client) RestoreDroplet(id, imageID int) error {
	_, err := c.send(DropletsEndpoint, fmt.Sprintf("%d/restore", id), Params{
		"image_id": imageID,
	})
	if err != nil {
		return err
	}
	return nil
}
