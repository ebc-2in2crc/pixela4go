package pixela

import "net/http"

// A Client manages communication with the Pixela User API.
type Client struct {
	UserName   string
	Token      string
	HTTPClient HTTPClient
}

// New return a new Client instance.
func New(userName, token string) *Client {
	return &Client{
		UserName:   userName,
		Token:      token,
		HTTPClient: &http.Client{},
	}
}

// User returns a new Pixela user API client.
func (c *Client) User() *User {
	return &User{UserName: c.UserName, Token: c.Token, httpClient: c.HTTPClient}
}

// UserProfile returns a new Pixela user profile API client.
func (c *Client) UserProfile() *UserProfile {
	return &UserProfile{UserName: c.UserName, Token: c.Token, httpClient: c.HTTPClient}
}

// Graph returns a new Pixela graph API client.
func (c *Client) Graph() *Graph {
	return &Graph{UserName: c.UserName, Token: c.Token, httpClient: c.HTTPClient}
}

// Pixel returns a new Pixela pixel API client.
func (c *Client) Pixel() *Pixel {
	return &Pixel{UserName: c.UserName, Token: c.Token, httpClient: c.HTTPClient}
}

// Webhook returns a new Pixela webhook API client.
func (c *Client) Webhook() *Webhook {
	return &Webhook{UserName: c.UserName, Token: c.Token, httpClient: c.HTTPClient}
}
