package ganboard

import "encoding/json"

// GetVersion https://docs.kanboard.org/en/latest/api/application_procedures.html#getversion
func (c *Client) GetVersion() (string, error) {
	query := request{
		Client: c,
		Method: "getVersion",
	}
	response, err := query.decodeString()
	return response, err
}

// GetTimezone https://docs.kanboard.org/en/latest/api/application_procedures.html#gettimezone
func (c *Client) GetTimezone() (string, error) {
	query := request{
		Client: c,
		Method: "getTimezone",
	}
	response, err := query.decodeString()
	return response, err
}

// GetDefaultTaskColors https://docs.kanboard.org/en/latest/api/application_procedures.html#getDefaultTaskColors
func (c *Client) GetDefaultTaskColors() (map[string]Color, error) {
	query := request{
		Client: c,
		Method: "getDefaultTaskColors",
	}
	response, err := query.decodeTaskColors()
	return response, err
}

// GetDefaultTaskColor https://docs.kanboard.org/en/latest/api/application_procedures.html#getDefaultTaskColor
func (c *Client) GetDefaultTaskColor() (string, error) {
	query := request{
		Client: c,
		Method: "getDefaultTaskColor",
	}
	response, err := query.decodeString()
	return response, err
}

// GetColorList https://docs.kanboard.org/en/latest/api/application_procedures.html#getColorList
func (c *Client) GetColorList() (map[string]string, error) {
	query := request{
		Client: c,
		Method: "getColorList",
	}
	response, err := query.decodeMapStringString()
	return response, err
}

// GetApplicationRoles https://docs.kanboard.org/en/latest/api/application_procedures.html#getApplicationRoles
func (c *Client) GetApplicationRoles() (map[string]string, error) {
	query := request{
		Client: c,
		Method: "getApplicationRoles",
	}
	response, err := query.decodeMapStringString()
	return response, err
}

// GetProjectRoles https://docs.kanboard.org/en/latest/api/application_procedures.html#getProjectRoles
func (c *Client) GetProjectRoles() (map[string]string, error) {
	query := request{
		Client: c,
		Method: "getProjectRoles",
	}
	response, err := query.decodeMapStringString()
	return response, err
}

// Color definition
type Color struct {
	Name       string `json:"name"`
	Background string `json:"background"`
	Border     string `json:"border"`
}

func (r *request) decodeTaskColors() (map[string]Color, error) {
	rsp, err := r.Client.Request(*r)
	if err != nil {
		return nil, err
	}

	body := struct {
		JSONRPC string           `json:"jsonrpc"`
		ID      int              `json:"id"`
		Result  map[string]Color `json:"result"`
	}{}

	err = json.NewDecoder(rsp.Body).Decode(&body)
	return body.Result, err
}
