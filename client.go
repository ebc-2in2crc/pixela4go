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

// Channel returns a new Pixela channel API client.
func (c *Client) Channel() *Channel {
	return &Channel{UserName: c.UserName, Token: c.Token}
}

// Graph returns a new Pixela graph API client.
func (c *Client) Graph() *Graph {
	return &Graph{UserName: c.UserName, Token: c.Token}
}

// Pixel returns a new Pixela pixel API client.
func (c *Client) Pixel() *Pixel {
	return &Pixel{UserName: c.UserName, Token: c.Token}
}

// Notification returns a new Pixela notification API client.
func (c *Client) Notification() *Notification {
	return &Notification{UserName: c.UserName, Token: c.Token}
}

// Webhook returns a new Pixela webhook API client.
func (c *Client) Webhook() *Webhook {
	return &Webhook{UserName: c.UserName, Token: c.Token}
}
