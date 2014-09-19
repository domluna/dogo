package dogo

// RebootDroplet reboots the droplet.
func (c *Client) RebootDroplet(id int) error {
	return c.DoAction(DropletEndpoint, id, Params{
		"type": "reboot",
	})
}

// PowerCycleDroplet powers cycles the droplet(power off then back on).
func (c *Client) PowerCycleDroplet(id int) error {
	return c.DoAction(DropletEndpoint, id, Params{
		"type": "power_cycle",
	})
}

// ShutdownDroplet powers off the droplet. This is a graceful way
// to shutdown the droplet.
func (c *Client) ShutdownDroplet(id int) error {
	return c.DoAction(DropletEndpoint, id, Params{
		"type": "shutdown",
	})
}

// PowerOffDroplet powers off the droplet. This is a hard shutdown
// and should only be used if a "shutdown" was not successful.
func (c *Client) PowerOffDroplet(id int) error {
	return c.DoAction(DropletEndpoint, id, Params{
		"type": "power_off",
	})
}

// PowerOnDroplet powers on the droplet.
func (c *Client) PowerOnDroplet(id int) error {
	return c.DoAction(DropletEndpoint, id, Params{
		"type": "power_on",
	})
}

// PasswordResetDroplet resets the password on the droplet.
func (c *Client) PasswordResetDroplet(id int) error {
	return c.DoAction(DropletEndpoint, id, Params{
		"type": "password_reset",
	})
}

// ResizeDroplet resizes the droplet to the new size.
func (c *Client) ResizeDroplet(id int, size string) error {
	return c.DoAction(DropletEndpoint, id, Params{
		"type": "resize",
		"size": size,
	})
}

// RestoreDroplet rebuilds the droplet using a backup image. The image ID
// passed must be a backup image of the current droplet. Leaves SSH keys
// intact.
func (c *Client) RestoreDroplet(id int, image int) error {
	return c.DoAction(DropletEndpoint, id, Params{
		"type":  "restore",
		"image": image,
	})
}

// RebuildDroplet rebuilds the droplet. Can use an image id
// or an image slug.
func (c *Client) RebuildDroplet(id int, image interface{}) error {
	return c.DoAction(DropletEndpoint, id, Params{
		"type":  "rebuild",
		"image": image,
	})
}

// RenameDroplet renames the droplet to the new name.
func (c *Client) RenameDroplet(id int, name string) error {
	return c.DoAction(DropletEndpoint, id, Params{
		"type": "rename",
		"name": name,
	})
}

// ChangeKernel changes the kernel of the droplet given the kernel id.
func (c *Client) ChangeKernel(id int, kernel int) error {
	return c.DoAction(DropletEndpoint, id, Params{
		"type":   "change_kernel",
		"kernel": kernel,
	})
}

// EnableIPV6 enables ipv6 on the droplet.
func (c *Client) EnableIPV6(id int) error {
	return c.DoAction(DropletEndpoint, id, Params{
		"type": "enable_ipv6",
	})
}

// DisableBackups disables backup images for the droplet.
func (c *Client) DisableBackups(id int) error {
	return c.DoAction(DropletEndpoint, id, Params{
		"type": "disable_backups",
	})
}

// EnablePrivateNetworking enables private networking on the droplet.
func (c *Client) EnablePrivateNetworking(id int) error {
	return c.DoAction(DropletEndpoint, id, Params{
		"type": "enable_private_networking",
	})
}

// SnapshotDroplet snapshots the droplet in its current state, creating an image.
func (c *Client) SnapshotDroplet(id int, name string) error {
	return c.DoAction(DropletEndpoint, id, Params{
		"type": "enable_private_networking",
		"name": name,
	})
}
