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
		Method:  "getGroupd",
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

// GetMemberGroups https://docs.kanboard.org/en/latest/api/group_member_procedures.html#getmembergroups
func (c *Client) GetMemberGroups(userID int) ([]Group, error) {
	request := request{
		JSONRPC: "2.0",
		Method:  "getMemberGroups",
		ID:      1,
		Params: map[string]int{
			"user_id": userID,
		},
	}

	rsp, err := c.Request(request)
	body := responseGroups{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// GetGroupMembers https://docs.kanboard.org/en/latest/api/group_member_procedures.html#getgroupmembers
func (c *Client) GetGroupMembers(groupID int) ([]User, error) {
	request := request{
		JSONRPC: "2.0",
		Method:  "getGroupMembers",
		ID:      1,
		Params: map[string]int{
			"group_id": groupID,
		},
	}

	rsp, err := c.Request(request)
	body := responseUsers{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// AddGroupMember https://docs.kanboard.org/en/latest/api/group_member_procedures.html#addgroupmember
func (c *Client) AddGroupMember(groupID int, userID int) (bool, error) {
	request := request{
		JSONRPC: "2.0",
		Method:  "addGroupMember",
		ID:      1,
		Params: map[string]int{
			"group_id": groupID,
			"user_id":  userID,
		},
	}

	rsp, err := c.Request(request)
	body := responseBoolean{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// RemoveGroupMember https://docs.kanboard.org/en/latest/api/group_member_procedures.html#removegroupmember
func (c *Client) RemoveGroupMember(groupID int, userID int) (bool, error) {
	request := request{
		JSONRPC: "2.0",
		Method:  "removeGroupMember",
		ID:      1,
		Params: map[string]int{
			"group_id": groupID,
			"user_id":  userID,
		},
	}

	rsp, err := c.Request(request)
	body := responseBoolean{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// IsGroupMember https://docs.kanboard.org/en/latest/api/group_member_procedures.html#isgroupmember
func (c *Client) IsGroupMember(groupID int, userID int) (bool, error) {
	request := request{
		JSONRPC: "2.0",
		Method:  "isGroupMember",
		ID:      1,
		Params: map[string]int{
			"group_id": groupID,
			"user_id":  userID,
		},
	}

	rsp, err := c.Request(request)
	body := responseBoolean{}
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
