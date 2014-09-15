/*
Package dogo provides an client to the DigitalOcean API.

The digitalocean package provides the REST API to interact with DigitalOcean. This is
used under the hood in the other packages which represent the digitalocean resource.

Current clients include:

	droplets
	sizes
	images
	domains
	actions
	regions
	keys

Each client requires a DigitalOcean token. Tokens can either be read, write or both read/write; so make sure you have the correct token permissions.

If you export the token as follows:

$ export $DIGITALOCEAN_TOKEN="token"

You could authenticate like so:

	import (
		"github.com/domluna/dogo/digitalocean"
	)

	func main() {
		token, err := digitalocean.EnvAuth()
		if err != nil {
			// make sure the env variable is set
		}
		// use token with client apis
	}

For example using the droplet client will go as follows:

	import (
		"github.com/domluna/dogo/droplet"
	)

	func main() {
		cli := droplet.NewClient("your token here")

		// get all droplets
		droplets, err := cli.GetAll()
		if err != nil {
			// deal with error
		}
		// do stuff with droplets
	}

All interactions to DigitalOcean take place within the client.
*/
package digitalocean
