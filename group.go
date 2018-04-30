package ganboard

import "encoding/json"

// CreateGroup https://docs.kanboard.org/en/latest/api/group_procedures.html#creategroup
func (c *Client) CreateGroup(params Group) (int, error) {
	request := request{
		JSONRPC: "2.0",
		Method:  "createGroup",
		ID:      1,
		Params:  params,
	}

	rsp, err := c.Request(request)
	body := responseInt{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// UpdateGroup https://docs.kanboard.org/en/latest/api/group_procedures.html#updategroup
func (c *Client) UpdateGroup(params Group) (bool, error) {
	request := request{
		JSONRPC: "2.0",
		Method:  "updateGroup",
		ID:      1,
		Params:  params,
	}

	rsp, err := c.Request(request)
	body := responseBoolean{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// RemoveGroup  https://docs.kanboard.org/en/latest/api/group_procedures.html#removegroup
func (c *Client) RemoveGroup(groupID int) (bool, error) {
	request := request{
		JSONRPC: "2.0",
		Method:  "removeGroup",
		ID:      1,
		Params: map[string]int{
			"group_id": groupID,
		},
	}

	rsp, err := c.Request(request)
	body := responseBoolean{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// GetGroup https://docs.kanboard.org/en/latest/api/group_procedures.html#getgroup
func (c *Client) GetGroup(groupID int) (Group, error) {
	request := request{
		JSONRPC: "2.0",
		Method:  "isActiveUser",
		ID:      1,
		Params: map[string]int{
			"group_id": groupID,
		},
	}

	rsp, err := c.Request(request)
	body := responseGroup{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// GetAllGroups https://docs.kanboard.org/en/latest/api/group_procedures.html#getallgroups
func (c *Client) GetAllGroups() ([]Group, error) {
	request := request{
		JSONRPC: "2.0",
		Method:  "getAllGroups",
		ID:      1,
	}

	rsp, err := c.Request(request)
	body := responseGroups{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// Group input for CreateGroup
type Group struct {
	ID         int    `json:"id,string,omitempty"`
	Name       string `json:"name,string,omitempty"`
	ExternalID string `json:"external_id,omitempty"`
}

type responseGroup struct {
	JSONRPC string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  Group  `json:"result"`
}

type responseGroups struct {
	JSONRPC string  `json:"jsonrpc"`
	ID      int     `json:"id"`
	Result  []Group `json:"result"`
}
