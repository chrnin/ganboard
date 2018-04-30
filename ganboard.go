package ganboard

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// Client access to kanboard API methods
type Client struct {
	Endpoint string
	Username string
	Password string
}

// Request sends jsonrpc request to kanboard and returns http.Response
func (c *Client) Request(request request) (*http.Response, error) {
	httpClient := &http.Client{}

	jsonrpc := new(bytes.Buffer)
	json.NewEncoder(jsonrpc).Encode(request)

	req, err := http.NewRequest(
		"POST",
		c.Endpoint,
		jsonrpc,
	)

	req.SetBasicAuth(c.Username, c.Password)
	rsp, err := httpClient.Do(req)
	return rsp, err
}

type response struct {
	JSONRPC string      `json:"jsonrpc"`
	ID      int         `json:"id"`
	Result  interface{} `json:"result"`
}

type responseInt struct {
	JSONRPC string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  int    `json:"result"`
}

type responseString struct {
	JSONRPC string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  string `json:"result"`
}

type responseBoolean struct {
	JSONRPC string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  bool   `json:"result"`
}

type responseFloat64 struct {
	JSONRPC string  `json:"jsonrpc"`
	ID      int     `json:"id"`
	Result  float64 `json:"result"`
}

type responseMapStringString struct {
	JSONRPC string            `json:"jsonrpc"`
	ID      int               `json:"id"`
	Result  map[string]string `json:"result"`
}

type request struct {
	JSONRPC string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	ID      int         `json:"id,string"`
	Params  interface{} `json:"params,omitempty"`
}
