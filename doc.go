/*
Package dogo provides an client to the DigitalOcean V2 API.

Current endpoints include:

        droplets
        images
        keys
        sizes
        regions
        domains
        actions

Each client requires a DigitalOcean token. Tokens can either be read, write or both read/write; so make sure you have the correct token permissions.

If you export the token as follows:

        $ export $DIGITALOCEAN_TOKEN="token"

For example using the droplet client will go as follows:

        package main

	import (
		"github.com/domluna/dogo"
	)

	func main() {

                // If the token is the empty string("") then it'll attempt
                // to fill the value under env var $DIGITALOCEAN_TOKEN
		client := dogo.NewClient("") // $DIGITALOCEAN_TOKEN

		// get all droplets
		droplets, err := client.ListDroplets()
		if err != nil {
			// deal with error
		}
		// do stuff with droplets
	}
*/
package dogo
