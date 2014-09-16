package digitalocean

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
