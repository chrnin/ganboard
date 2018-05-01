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

// Request sends jsonrpc request to kanboard and returns *http.Response.
// FIXME set incremental id to track requests
// FIXME implement batch requests
func (c *Client) Request(request request) (*http.Response, error) {
	// construct request
	httpClient := &http.Client{}
	request.JSONRPC = "2.0"
	request.ID = 1
	jsonrpc := new(bytes.Buffer)
	json.NewEncoder(jsonrpc).Encode(request)
	req, err := http.NewRequest(
		"POST",
		c.Endpoint,
		jsonrpc,
	)

	if err != nil {
		rsp := new(http.Response)
		return rsp, err
	}

	// set auth and send to kanboard
	req.SetBasicAuth(c.Username, c.Password)
	response, err := httpClient.Do(req)

	return response, err
}

type response struct {
	JSONRPC string      `json:"jsonrpc"`
	ID      int         `json:"id"`
	Result  interface{} `json:"result"`
}

type request struct {
	JSONRPC string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	ID      int         `json:"id,string"`
	Params  interface{} `json:"params,omitempty"`
	Client  *Client     `json:"-"`
}

func (r *request) decodeInt() (int, error) {
	rsp, err := r.Client.Request(*r)
	if err != nil {
		return 0, err
	}

	body := struct {
		JSONRPC string `json:"jsonrpc"`
		ID      int    `json:"id"`
		Result  int    `json:"result,string"`
	}{}
	err = json.NewDecoder(rsp.Body).Decode(&body)
	return body.Result, err
}

func (r *request) decodeInterface() (interface{}, error) {
	rsp, err := r.Client.Request(*r)
	if err != nil {
		return nil, err
	}

	body := struct {
		JSONRPC string      `json:"jsonrpc"`
		ID      int         `json:"id"`
		Result  interface{} `json:"result"`
	}{}
	err = json.NewDecoder(rsp.Body).Decode(&body)
	return body.Result, err
}

func (r *request) decodeString() (string, error) {
	rsp, err := r.Client.Request(*r)

	if err != nil {
		return "", err
	}

	body := struct {
		JSONRPC string `json:"jsonrpc"`
		ID      int    `json:"id"`
		Result  string `json:"result"`
	}{}

	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

func (r *request) decodeBoolean() (bool, error) {
	rsp, err := r.Client.Request(*r)

	if err != nil {
		return false, err
	}

	body := struct {
		JSONRPC string `json:"jsonrpc"`
		ID      int    `json:"id"`
		Result  bool   `json:"result"`
	}{}

	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

func (r *request) decodeFloat64() (float64, error) {
	rsp, err := r.Client.Request(*r)

	if err != nil {
		return 0, err
	}

	body := struct {
		JSONRPC string  `json:"jsonrpc"`
		ID      int     `json:"id"`
		Result  float64 `json:"result"`
	}{}

	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

func (r *request) decodeMapStringString() (map[string]string, error) {
	rsp, err := r.Client.Request(*r)

	if err != nil {
		return nil, err
	}

	body := struct {
		JSONRPC string            `json:"jsonrpc"`
		ID      int               `json:"id"`
		Result  map[string]string `json:"result"`
	}{}

	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

func (r *request) decodeMapIntString() (map[int]string, error) {
	rsp, err := r.Client.Request(*r)

	if err != nil {
		return nil, err
	}

	body := struct {
		JSONRPC string         `json:"jsonrpc"`
		ID      int            `json:"id"`
		Result  map[int]string `json:"result"`
	}{}

	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}
