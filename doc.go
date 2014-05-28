/*
Package dogo provides an client to the DigitalOcean API.

Creation of a Client:

	// CLIENT_ID and API_KEY is your DigitalOcean
	// client id and api key.

	// See https://cloud.digitalocean.com/api_access.
	auth := dogo.Auth{CLIENT_ID, API_KEY}

	// Make a client with credentials
	client := dogo.NewClient(auth)


A Client provides access to DigitalOcean endpoints such as
Droplets, Images, Regions, SSH Keys, Sizes, etc.

To see all your droplets:

	droplets, err := client.GetDroplets()
	...
	// do stuff with droplets

The API is similar for other service, for example: to get all
available regions:

	regions, err := client.GetRegions()
	...
	// do stuff with regions

To create a new Droplet:

	// Creates a new droplet named "my_droplet"
	// SizeID: 66 => "512MB"
	// ImageID: 3668014 => "Docker Image"
	// RegionID: 4 => "nyc2"
	// BackupsActive enables backups
	//
	// 1234 is a ssh key id, this can be a list of ints
	// for multiple ids.
	//
	// The last argument sets whether you want a private network id.
	d := dogo.Droplet{
		Name: "droppy",
		SizeID: 66,
		ImageID: 3668014,
		RegionID: 4,
		BackupsActive: true,
	}
	droplet, err := client.CreateDroplet(d, 1234, true)
	...

To add a ssh key to your account:

	// public key is the byte content of a ssh key file
	key, err := client.AddSSHKey("super_secret_key", public_key)
	...

The region, image and sizes apis are very simple, for the most just
one maybe two endpoints.

The droplet and ssh key apis are where most of the interaction will be.
*/
package dogo
