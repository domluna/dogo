package dogo

// TransferImage transfers an image to another region.
func (c *Client) TransferImage(id int, region string) error {
	return c.DoAction(ImageEndpoint, id, Params{
		"type": "transfer",
		"region": region,
	})
}
