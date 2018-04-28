package ganboard

import (
	"io"
	"net/http"
)

// Client access to kanboard API methods
type Client struct {
	Endpoint string
	Username string
	Password string
}

// Request sends jsonrpc request to kanboard and returns http.Response
func (c *Client) Request(jsonrpc io.Reader) (*http.Response, error) {
	httpClient := &http.Client{}
	req, err := http.NewRequest(
		"POST",
		c.Endpoint,
		jsonrpc,
	)

	req.SetBasicAuth(c.Username, c.Password)
	rsp, err := httpClient.Do(req)
	return rsp, err
}
