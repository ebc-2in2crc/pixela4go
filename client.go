package pixela

// A Client manages communication with the Pixela User API.
type Client struct {
	UserName string
	Token    string
}

// New return a new Client instance.
func New(userName, token string) *Client {
	return &Client{UserName: userName, Token: token}
}

// User returns a new Pixela user API client.
func (c *Client) User() *User {
	return &User{UserName: c.UserName, Token: c.Token}
}

// Graph returns a new Pixela graph API client.
func (c *Client) Graph() *Graph {
	return &Graph{UserName: c.UserName, Token: c.Token}
}

// Pixel returns a new Pixela pixel API client.
func (c *Client) Pixel() *Pixel {
	return &Pixel{UserName: c.UserName, Token: c.Token}
}

// Webhook returns a new Pixela webhook API client.
func (c *Client) Webhook() *Webhook {
	return &Webhook{UserName: c.UserName, Token: c.Token}
}
