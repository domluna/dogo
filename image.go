package dogo

import (
	"fmt"
)

// Kernel is a DigitalOcean Kernel.
type Kernel struct {
	ID      int    `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Version string `json:"version,omitempty"`
}

type Kernels []*Kernel

// Snapshot is a DigitalOcean snapshot.
type Snapshot struct {
	ID           int      `json:"id,omitempty"`
	Name         string   `json:"name,omitempty"`
	Distribution string   `json:"distribution,omitempty"`
	Slug         string   `json:"slug,omitempty,omitempty"`
	Public       bool     `json:"public,omitempty"`
	Regions      []string `json:"regions,omitempty"`
	ActionIDs    []int    `json:"action_ids,omitempty"`
}

type Snapshots []*Snapshot

// Backup is a DigitalOcean backup.
type Backup struct {
	ID           int      `json:"id,omitempty"`
	Name         string   `json:"name,omitempty"`
	Distribution string   `json:"distribution,omitempty"`
	Slug         string   `json:"slug,omitempty,omitempty"`
	Public       bool     `json:"public,omitempty"`
	Regions      []string `json:"regions,omitempty"`
	ActionIDs    []int    `json:"action_ids,omitempty"`
}

type Backups []*Backup

// Representation for a DigitalOcean Image.
type Image struct {
	ID           int      `json:"id,omitempty"`
	Name         string   `json:"name,omitempty"`
	Distribution string   `json:"distribution,omitempty"`
	Slug         string   `json:"slug,omitempty,omitempty"`
	Public       bool     `json:"public,omitempty"`
	Regions      []string `json:"regions,omitempty"`
	CreatedAt    string   `json:"created_at,omitempty"`
}

type Images []*Image

// UpdateImageOpts contains options used when updating a image.
type UpdateImageOpts struct {
	Name string `json:"name"`
}

// ListImages retrieves all images on your account.
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

// GetImage retrieves an image by its id or slug.
func (c *Client) GetImage(v interface{}) (*Image, error) {
	u := fmt.Sprintf("%s/%v", ImageEndpoint, v)
	s := struct {
		Image `json:"image,omitempty"`
	}{}
	err := c.get(u, &s)
	if err != nil {
		return nil, err
	}
	return &s.Image, nil
}

// DeleteImage deletes an image given its id.
func (c *Client) DeleteImage(id int) error {
	u := fmt.Sprintf("%s/%d", ImageEndpoint, id)
	err := c.delete(u)
	if err != nil {
		return err
	}
	return nil
}

// UpdateImage updates the image's name given its id.
func (c *Client) UpdateImage(id int, opts *UpdateImageOpts) (*Image, error) {
	u := fmt.Sprintf("%s/%d", ImageEndpoint, id)
	s := struct {
		Image `json:"image,omitempty"`
	}{}
	err := c.put(u, opts, &s)
	if err != nil {
		return nil, err
	}
	return &s.Image, nil
}
