package dogo

import "fmt"

// Droplet respresents a DigitalOcean droplet.
type Droplet struct {
	ID          int      `json:"id,omitempty"`
	Name        string   `json:"name,omitempty"`
	Region      Region   `json:"region,omitempty"`
	Image       Image    `json:"image,omitempty"`
	Kernel      Kernel   `json:"kernel,omitempty"`
	Size        Size     `json:"size,omitempty"`
	Locked      bool     `json:"locked,omitempty"`
	Status      string   `json:"status,omitempty"`
	Networks    Networks `json:"networks,omitempty"`
	BackupIDs   []int    `json:"backups_ids,omitempty"`
	SnapshotIDs []int    `json:"snapshot_ids,omitempty"`
	ActionIDs   []int    `json:"action_ids,omitempty"`
}

type Droplets []Droplet

type DropletClient struct {
	Client
}

// GetDroplets returns all users droplets, active or otherwise.
func (dc *DropletClient) GetDroplets() (Droplets, error) {
	s := struct {
		Droplets `json:"droplets,omitempty"`
		Meta     `json:"meta,omitempty"`
	}{}
	err := dc.Client.GetX(DropletsEndpoint, &s)
	if err != nil {
		return nil, err
	}
	return s.Droplets, nil
}

// GetDroplet return an individual droplet based on the passed id.
func (dc *DropletClient) GetDroplet(id int) (Droplet, error) {
	u := fmt.Sprintf("%s/%d", DropletsEndpoint, id)
	s := struct {
		Droplet `json:"droplet,omitempty"`
	}{}

	err := dc.Client.GetX(u, &s)
	if err != nil {
		return s.Droplet, err
	}
	return s.Droplet, nil
}

// CreateDroplet creates a droplet based on based specs.
func (dc *DropletClient) CreateDroplet(opts map[string]interface{}) (Droplet, error) {
	var d Droplet
	req, err := dc.Client.Post(DropletsEndpoint, opts)
	if err != nil {
		return d, err
	}

	err = dc.Client.DoRequest(req, &d)
	if err != nil {
		return d, err
	}
	return d, nil
}

// DestroyDroplet destroys a droplet. CAUTION - this is irreversible.
// There may be more appropriate options.
func (dc *DropletClient) DestroyDroplet(id int) error {
	u := fmt.Sprintf("%s/%d", DropletsEndpoint, id)
	req, err := dc.Client.Del(u)
	if err != nil {
		return err
	}

	err = dc.Client.DoRequest(req, nil)
	if err != nil {
		return err
	}
	return nil
}

// ResizeDroplet droplet resizes a droplet. Sizes are based on
// the digitalocean sizes api.
func (dc *DropletClient) ResizeDroplet(id int, size string) error {
	u := fmt.Sprintf("%s/%d/actions", DropletsEndpoint, id)
	payload := map[string]interface{}{
		"type": "resize",
		"size": size,
	}

	req, err := dc.Client.Post(u, payload)
	if err != nil {
		return err
	}

	err = dc.Client.DoRequest(req, nil)
	if err != nil {
		return err
	}
	return nil
}

// RebootDroplet reboots the a droplet. This is the preferred method
// to use if a server is not responding.
func (dc *DropletClient) RebootDroplet(id int) error {
	u := fmt.Sprintf("%s/%d/actions", DropletsEndpoint, id)
	payload := map[string]interface{}{
		"type": "reboot",
	}

	req, err := dc.Client.Post(u, payload)
	if err != nil {
		return err
	}

	err = dc.Client.DoRequest(req, nil)
	if err != nil {
		return err
	}
	return nil
}

// RebootDroplet rebuilds a droplet with a default image. This can be
// useful if you want to use a different image but keep the ip address
// of the droplet.
func (dc *DropletClient) RebuildDroplet(id, imageID int) error {
	u := fmt.Sprintf("%s/%d/actions", DropletsEndpoint, id)
	payload := map[string]interface{}{
		"type":  "rebuild",
		"image": imageID,
	}

	req, err := dc.Client.Post(u, payload)
	if err != nil {
		return err
	}

	err = dc.Client.DoRequest(req, nil)
	if err != nil {
		return err
	}
	return nil
}

// StopDroplet powers off a running droplet, the droplet will remain
// in your account.
func (dc *DropletClient) PowerOffDroplet(id int) error {
	u := fmt.Sprintf("%s/%d/actions", DropletsEndpoint, id)
	payload := map[string]interface{}{
		"type": "power_off",
	}

	req, err := dc.Client.Post(u, payload)
	if err != nil {
		return err
	}

	err = dc.Client.DoRequest(req, nil)
	if err != nil {
		return err
	}
	return nil
}

// StartDroplet powers on a powered off droplet.
func (dc *DropletClient) PowerOnDroplet(id int) error {
	u := fmt.Sprintf("%s/%d/actions", DropletsEndpoint, id)
	payload := map[string]interface{}{
		"type": "power_on",
	}

	req, err := dc.Client.Post(u, payload)
	if err != nil {
		return err
	}

	err = dc.Client.DoRequest(req, nil)
	if err != nil {
		return err
	}
	return nil
}

// RestoreDroplet allows you to restore a droplet to a previous image
// or snapshot. This will be a mirror copy of the image or snapshot to
// your droplet.
func (dc *DropletClient) RestoreDroplet(id, imageID int) error {
	u := fmt.Sprintf("%s/%d/actions", DropletsEndpoint, id)
	payload := map[string]interface{}{
		"type":  "restore",
		"image": imageID,
	}

	req, err := dc.Client.Post(u, payload)
	if err != nil {
		return err
	}

	err = dc.Client.DoRequest(req, nil)
	if err != nil {
		return err
	}
	return nil
}
