package dogo

import "fmt"

// Droplet respresents a DigitalOcean droplet.
type Droplet struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Region      Region `json:"region,omitempty"`
	Image       Image  `json:"image,omitempty"`
	Kernel      Kernel `json:"kernel,omitempty"`
	Size        Size   `json:"size,omitempty"`
	Locked      bool   `json:"locked,omitempty"`
	Status      string `json:"status,omitempty"`
	Networks    string `json:"networks,omitempty"`
	BackupIDs   []int  `json:"backups_ids,omitempty"`
	SnapshotIDs []int  `json:"snapshot_ids,omitempty"`
	ActionIDs   []int  `json:"action_ids,omitempty"`
}

type Droplets []Droplet

type DropletClient struct {
	Client
}

// GetDroplets returns all users droplets, active or otherwise.
func (c *Client) GetDroplets() (Droplets, error) {
	var d Droplets
	err := c.Get(DropletsEndpoint, &d)

	if err != nil {
		return nil, err
	}
	return d, nil
}

// GetDroplet return an individual droplet based on the passed id.
func (c *Client) GetDroplet(id int) (Droplet, error) {
	u := fmt.Sprintf("%s/%d", DropletsEndpoint, id)
	var d Droplet
	err := c.Get(u, &d)

	if err != nil {
		return d, err
	}
	return d, nil
}

// CreateDroplet creates a droplet based on based specs.
func (c *Client) CreateDroplet(v map[string]interface{}) (Droplet, error) {
	var d Droplet
	err := c.Post(DropletsEndpoint, v, &d)

	if err != nil {
		return d, err
	}
	return d, nil
}

// DestroyDroplet destroys a droplet. CAUTION - this is irreversible.
// There may be more appropriate options.
func (c *Client) DestroyDroplet(id int) error {
	u := fmt.Sprintf("%s/%d", DropletsEndpoint, id)
	err := c.Del(u)

	if err != nil {
		return err
	}
	return nil
}

// ResizeDroplet droplet resizes a droplet. Sizes are based on
// the digitalocean sizes api.
func (c *Client) ResizeDroplet(id int, size string) error {
	u := fmt.Sprintf("%s/%d/actions", DropletsEndpoint, id)
	payload := map[string]interface{}{
		"type": "resize",
		"size": size,
	}
	err := c.Post(u, payload, nil)
	if err != nil {
		return err
	}
	return nil
}

// RebootDroplet reboots the a droplet. This is the preferred method
// to use if a server is not responding.
func (c *Client) RebootDroplet(id int) error {
	u := fmt.Sprintf("%s/%d/actions", DropletsEndpoint, id)
	payload := map[string]interface{}{
		"type": "reboot",
	}
	err := c.Post(u, payload, nil)
	if err != nil {
		return err
	}
	return nil
}

// RebootDroplet rebuilds a droplet with a default image. This can be
// useful if you want to use a different image but keep the ip address
// of the droplet.
func (c *Client) RebuildDroplet(id, imageID int) error {
	u := fmt.Sprintf("%s/%d/actions", DropletsEndpoint, id)
	payload := map[string]interface{}{
		"type":  "rebuild",
		"image": imageID,
	}
	err := c.Post(u, payload, nil)
	if err != nil {
		return err
	}
	return nil
}

// StopDroplet powers off a running droplet, the droplet will remain
// in your account.
func (c *Client) PowerOffDroplet(id int) error {
	u := fmt.Sprintf("%s/%d/actions", DropletsEndpoint, id)
	payload := map[string]interface{}{
		"type": "power_off",
	}
	err := c.Post(u, payload, nil)
	if err != nil {
		return err
	}
	return nil
}

// StartDroplet powers on a powered off droplet.
func (c *Client) PowerOnDroplet(id int) error {
	u := fmt.Sprintf("%s/%d/actions", DropletsEndpoint, id)
	payload := map[string]interface{}{
		"type": "power_on",
	}
	err := c.Post(u, payload, nil)
	if err != nil {
		return err
	}
	return nil
}

// RestoreDroplet allows you to restore a droplet to a previous image
// or snapshot. This will be a mirror copy of the image or snapshot to
// your droplet.
func (c *Client) RestoreDroplet(id, imageID int) error {
	u := fmt.Sprintf("%s/%d/actions", DropletsEndpoint, id)
	payload := map[string]interface{}{
		"type":  "restore",
		"image": imageID,
	}
	err := c.Post(u, payload, nil)
	if err != nil {
		return err
	}
	return nil
}
