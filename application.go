package ganboard

import (
	"encoding/json"
)

// GetVersion https://docs.kanboard.org/en/latest/api/application_procedures.html#getversion
func (c *Client) GetVersion() (string, error) {
	request := request{
		JSONRPC: "2.0",
		Method:  "getTimezone",
		ID:      "1",
	}

	rsp, err := c.Request(request)
	body := response{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result.(string), err
}

// GetTimezone https://docs.kanboard.org/en/latest/api/application_procedures.html#gettimezone
func (c *Client) GetTimezone() (string, error) {
	request := request{
		JSONRPC: "2.0",
		Method:  "getVersion",
		ID:      "1",
	}

	rsp, err := c.Request(request)
	body := response{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result.(string), err
}

// GetDefaultTaskColors https://docs.kanboard.org/en/latest/api/application_procedures.html#getDefaultTaskColors
func (c *Client) GetDefaultTaskColors() (map[string]Color, error) {
	request := request{
		JSONRPC: "2.0",
		Method:  "getDefaultTaskColors",
		ID:      "1",
	}

	rsp, err := c.Request(request)
	body := responseDefaultTaskColors{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// GetDefaultTaskColor https://docs.kanboard.org/en/latest/api/application_procedures.html#getDefaultTaskColor
func (c *Client) GetDefaultTaskColor() (string, error) {
	request := request{
		JSONRPC: "2.0",
		Method:  "getDefaultTaskColor",
		ID:      "1",
	}

	rsp, err := c.Request(request)
	body := response{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result.(string), err
}

// GetColorList https://docs.kanboard.org/en/latest/api/application_procedures.html#getColorList
func (c *Client) GetColorList() (map[string]string, error) {
	request := request{
		JSONRPC: "2.0",
		Method:  "getColorList",
		ID:      "1",
	}

	rsp, err := c.Request(request)
	body := responseColorList{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// GetApplicationRoles https://docs.kanboard.org/en/latest/api/application_procedures.html#getApplicationRoles
func (c *Client) GetApplicationRoles() (map[string]string, error) {
	request := request{
		JSONRPC: "2.0",
		Method:  "getApplicationRoles",
		ID:      "1",
	}

	rsp, err := c.Request(request)
	body := responseRolesMap{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// GetProjectRoles https://docs.kanboard.org/en/latest/api/application_procedures.html#getProjectRoles
func (c *Client) GetProjectRoles() (map[string]string, error) {
	request := request{
		JSONRPC: "2.0",
		Method:  "getProjectRoles",
		ID:      "1",
	}

	rsp, err := c.Request(request)
	body := responseRolesMap{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

type responseDefaultTaskColors struct {
	JSONRPC string           `json:"jsonrpc"`
	ID      int              `json:"id"`
	Result  map[string]Color `json:"result"`
}

type responseColorList struct {
	JSONRPC string            `json:"jsonrpc"`
	ID      int               `json:"id"`
	Result  map[string]string `json:"result"`
}

type responseRolesMap struct {
	JSONRPC string            `json:"jsonrpc"`
	ID      int               `json:"id"`
	Result  map[string]string `json:"result"`
}

// Color definition
type Color struct {
	Name       string `json:"name"`
	Background string `json:"background"`
	Border     string `json:"border"`
}
