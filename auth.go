package dogo

type Auth struct {
	ClientID string
	APIKey   string
}

type Client struct {
	Auth
}
