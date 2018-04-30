package ganboard

import "encoding/json"

// GetProjectUsers https://docs.kanboard.org/en/latest/api/project_permission_procedures.html#getprojectusers
func (c *Client) GetProjectUsers(projectID int) (map[int]string, error) {
	request := request{
		JSONRPC: "2.0",
		Method:  "getProjectUsers",
		ID:      1,
		Params: map[string]int{
			"project_id": projectID,
		},
	}

	rsp, err := c.Request(request)
	body := responseUserList{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// GetAssignableUsers https://docs.kanboard.org/en/latest/api/project_permission_procedures.html#getassignableusers
func (c *Client) GetAssignableUsers(params AssignableUsersParams) (map[int]string, error) {
	request := request{
		JSONRPC: "2.0",
		Method:  "getAssignableUsers",
		ID:      1,
		Params:  params,
	}

	rsp, err := c.Request(request)
	body := responseUserList{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// AddProjectUser https://docs.kanboard.org/en/latest/api/project_permission_procedures.html#addprojectuser
func (c *Client) AddProjectUser(params ProjectUserParams) (bool, error) {
	request := request{
		JSONRPC: "2.0",
		Method:  "addProjectUser",
		ID:      1,
		Params:  params,
	}

	rsp, err := c.Request(request)
	body := responseBoolean{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// AddProjectGroup https://docs.kanboard.org/en/latest/api/project_permission_procedures.html#addprojectgroup
func (c *Client) AddProjectGroup(params ProjectGroupParams) (bool, error) {
	request := request{
		JSONRPC: "2.0",
		Method:  "addProjectGroup",
		ID:      1,
		Params:  params,
	}

	rsp, err := c.Request(request)
	body := responseBoolean{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// RemoveProjectUser https://docs.kanboard.org/en/latest/api/project_permission_procedures.html#removeprojectuser
func (c *Client) RemoveProjectUser(params ProjectUserParams) (bool, error) {
	request := request{
		JSONRPC: "2.0",
		Method:  "removeProjectUser",
		ID:      1,
		Params:  params,
	}

	rsp, err := c.Request(request)
	body := responseBoolean{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// RemoveProjectGroup https://docs.kanboard.org/en/latest/api/project_permission_procedures.html#removeprojectgroup
func (c *Client) RemoveProjectGroup(params ProjectGroupParams) (bool, error) {
	request := request{
		JSONRPC: "2.0",
		Method:  "removeProjectGroup",
		ID:      1,
		Params:  params,
	}

	rsp, err := c.Request(request)
	body := responseBoolean{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// ChangeProjectUserRole https://docs.kanboard.org/en/latest/api/project_permission_procedures.html#changeprojectuserrole
func (c *Client) ChangeProjectUserRole(params ProjectUserParams) (bool, error) {
	request := request{
		JSONRPC: "2.0",
		Method:  "changeProjectUserRole",
		ID:      1,
		Params:  params,
	}

	rsp, err := c.Request(request)
	body := responseBoolean{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// ChangeProjectGroupRole https://docs.kanboard.org/en/latest/api/project_permission_procedures.html#changeprojectgrouprole
func (c *Client) ChangeProjectGroupRole(params ProjectUserParams) (bool, error) {
	request := request{
		JSONRPC: "2.0",
		Method:  "changeProjectGroupRole",
		ID:      1,
		Params:  params,
	}

	rsp, err := c.Request(request)
	body := responseBoolean{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// GetProjectUserRole https://docs.kanboard.org/en/latest/api/project_permission_procedures.html#getprojectuserrole
func (c *Client) GetProjectUserRole(projectID int, userID int) (string, error) {
	request := request{
		JSONRPC: "2.0",
		Method:  "getProjectUserRole",
		ID:      1,
		Params: map[string]int{
			"project_id": projectID,
			"user_id":    userID,
		},
	}

	rsp, err := c.Request(request)
	body := responseString{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

type responseUserList struct {
	JSONRPC string         `json:"jsonrpc"`
	ID      int            `json:"id"`
	Result  map[int]string `json:"result"`
}

// AssignableUsersParams input for GetAssignableUsers
type AssignableUsersParams struct {
	ProjectID          int  `json:"project_id"`
	PreprendUnassigned bool `json:"preprend_unasssigned,omitempty"`
}

// ProjectUserParams input for AddProjectUser
type ProjectUserParams struct {
	ProjectID int    `json:"project_id"`
	UserID    int    `json:"user_id"`
	Role      string `json:"role,omitempty"`
}

// ProjectGroupParams input for AddProjectGroup
type ProjectGroupParams struct {
	ProjectID int    `json:"project_id"`
	GroupID   int    `json:"group_id"`
	Role      string `json:"role,omitempty"`
}
