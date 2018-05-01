package ganboard

import "encoding/json"

// CreateGroup https://docs.kanboard.org/en/latest/api/group_procedures.html#creategroup
func (c *Client) CreateGroup(params Group) (int, error) {
	query := request{
		Client: c,
		Method: "createGroup",
		Params: params,
	}
	response, err := query.decodeInt()
	return response, err
}

// UpdateGroup https://docs.kanboard.org/en/latest/api/group_procedures.html#updategroup
func (c *Client) UpdateGroup(params Group) (bool, error) {
	query := request{
		Client: c,
		Method: "updateGroup",
		Params: params,
	}
	response, err := query.decodeBoolean()
	return response, err
}

// RemoveGroup  https://docs.kanboard.org/en/latest/api/group_procedures.html#removegroup
func (c *Client) RemoveGroup(groupID int) (bool, error) {
	query := request{
		Client: c,
		Method: "removeGroup",
		Params: map[string]int{
			"group_id": groupID,
		},
	}
	response, err := query.decodeBoolean()
	return response, err
}

// GetGroup https://docs.kanboard.org/en/latest/api/group_procedures.html#getgroup
func (c *Client) GetGroup(groupID int) (Group, error) {
	query := request{
		Client: c,
		Method: "getGroupd",
		Params: map[string]int{
			"group_id": groupID,
		},
	}
	response, err := query.decodeGroup()
	return response, err
}

// GetAllGroups https://docs.kanboard.org/en/latest/api/group_procedures.html#getallgroups
func (c *Client) GetAllGroups() ([]Group, error) {
	query := request{
		Client: c,
		Method: "getAllGroups",
	}
	response, err := query.decodeGroups()
	return response, err
}

// GetMemberGroups https://docs.kanboard.org/en/latest/api/group_member_procedures.html#getmembergroups
func (c *Client) GetMemberGroups(userID int) ([]Group, error) {
	query := request{
		Client: c,
		Method: "getMemberGroups",
		Params: map[string]int{
			"user_id": userID,
		},
	}
	response, err := query.decodeGroups()
	return response, err
}

// GetGroupMembers https://docs.kanboard.org/en/latest/api/group_member_procedures.html#getgroupmembers
func (c *Client) GetGroupMembers(groupID int) ([]User, error) {
	query := request{
		Client: c,
		Method: "getGroupMembers",
		Params: map[string]int{
			"group_id": groupID,
		},
	}
	response, err := query.decodeUsers()
	return response, err
}

// AddGroupMember https://docs.kanboard.org/en/latest/api/group_member_procedures.html#addgroupmember
func (c *Client) AddGroupMember(groupID int, userID int) (bool, error) {
	query := request{
		Client: c,
		Method: "addGroupMember",
		Params: map[string]int{
			"group_id": groupID,
			"user_id":  userID,
		},
	}
	response, err := query.decodeBoolean()
	return response, err
}

// RemoveGroupMember https://docs.kanboard.org/en/latest/api/group_member_procedures.html#removegroupmember
func (c *Client) RemoveGroupMember(groupID int, userID int) (bool, error) {
	query := request{
		Client: c,
		Method: "removeGroupMember",
		Params: map[string]int{
			"group_id": groupID,
			"user_id":  userID,
		},
	}
	response, err := query.decodeBoolean()
	return response, err
}

// IsGroupMember https://docs.kanboard.org/en/latest/api/group_member_procedures.html#isgroupmember
func (c *Client) IsGroupMember(groupID int, userID int) (bool, error) {
	query := request{
		Client: c,
		Method: "isGroupMember",
		Params: map[string]int{
			"group_id": groupID,
			"user_id":  userID,
		},
	}
	response, err := query.decodeBoolean()
	return response, err
}

// Group input for CreateGroup
type Group struct {
	ID         int    `json:"id,string,omitempty"`
	Name       string `json:"name,string,omitempty"`
	ExternalID string `json:"external_id,omitempty"`
}

func (r *request) decodeGroup() (Group, error) {
	rsp, err := r.Client.Request(*r)
	if err != nil {
		return Group{}, err
	}

	body := struct {
		JSONRPC string `json:"jsonrpc"`
		ID      int    `json:"id"`
		Result  Group  `json:"result"`
	}{}

	err = json.NewDecoder(rsp.Body).Decode(&body)
	return body.Result, err
}

func (r *request) decodeGroups() ([]Group, error) {
	rsp, err := r.Client.Request(*r)
	if err != nil {
		return nil, err
	}

	body := struct {
		JSONRPC string  `json:"jsonrpc"`
		ID      int     `json:"id"`
		Result  []Group `json:"result"`
	}{}

	err = json.NewDecoder(rsp.Body).Decode(&body)
	return body.Result, err
}
