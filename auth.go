package dogo

// Auth contains data required to authenticate 
// get DigitalOcean api.
type Auth struct {
	ClientID string
	APIKey   string
}

// Client is a wrapper around Auth, Clients are used
// to query the api.
// 
// To make a new Client call NewClient.
type Client struct {
	Auth
}

// NewClient creates a new Client.
func NewClient(auth Auth) *Client {
	return &Client{auth}
}
