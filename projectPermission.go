package ganboard

// GetProjectUsers https://docs.kanboard.org/en/latest/api/project_permission_procedures.html#getprojectusers
func (c *Client) GetProjectUsers(projectID int) (map[int]string, error) {
	query := request{
		Client: c,
		Method: "getProjectUsers",
		Params: map[string]int{
			"project_id": projectID,
		},
	}
	response, err := query.decodeMapIntString()
	return response, err
}

// GetAssignableUsers https://docs.kanboard.org/en/latest/api/project_permission_procedures.html#getassignableusers
func (c *Client) GetAssignableUsers(params AssignableUsersParams) (map[int]string, error) {
	query := request{
		Client: c,
		Method: "getAssignableUsers",
		Params: params,
	}
	response, err := query.decodeMapIntString()
	return response, err
}

// AddProjectUser https://docs.kanboard.org/en/latest/api/project_permission_procedures.html#addprojectuser
func (c *Client) AddProjectUser(params ProjectUserParams) (bool, error) {
	query := request{
		Client: c,
		Method: "addProjectUser",
		Params: params,
	}
	response, err := query.decodeBoolean()
	return response, err
}

// AddProjectGroup https://docs.kanboard.org/en/latest/api/project_permission_procedures.html#addprojectgroup
func (c *Client) AddProjectGroup(params ProjectGroupParams) (bool, error) {
	query := request{
		Client: c,
		Method: "addProjectGroup",
		Params: params,
	}
	response, err := query.decodeBoolean()
	return response, err
}

// RemoveProjectUser https://docs.kanboard.org/en/latest/api/project_permission_procedures.html#removeprojectuser
func (c *Client) RemoveProjectUser(params ProjectUserParams) (bool, error) {
	query := request{
		Client: c,
		Method: "removeProjectUser",
		Params: params,
	}
	response, err := query.decodeBoolean()
	return response, err
}

// RemoveProjectGroup https://docs.kanboard.org/en/latest/api/project_permission_procedures.html#removeprojectgroup
func (c *Client) RemoveProjectGroup(params ProjectGroupParams) (bool, error) {
	query := request{
		Client: c,
		Method: "removeProjectGroup",
		ID:     1,
		Params: params,
	}
	response, err := query.decodeBoolean()
	return response, err
}

// ChangeProjectUserRole https://docs.kanboard.org/en/latest/api/project_permission_procedures.html#changeprojectuserrole
func (c *Client) ChangeProjectUserRole(params ProjectUserParams) (bool, error) {
	query := request{
		Client: c,
		Method: "changeProjectUserRole",
		Params: params,
	}
	response, err := query.decodeBoolean()
	return response, err
}

// ChangeProjectGroupRole https://docs.kanboard.org/en/latest/api/project_permission_procedures.html#changeprojectgrouprole
func (c *Client) ChangeProjectGroupRole(params ProjectUserParams) (bool, error) {
	query := request{
		Client: c,
		Method: "changeProjectGroupRole",
		Params: params,
	}
	response, err := query.decodeBoolean()
	return response, err
}

// GetProjectUserRole https://docs.kanboard.org/en/latest/api/project_permission_procedures.html#getprojectuserrole
func (c *Client) GetProjectUserRole(projectID int, userID int) (string, error) {
	query := request{
		Client: c,
		Method: "getProjectUserRole",
		Params: map[string]int{
			"project_id": projectID,
			"user_id":    userID,
		},
	}
	response, err := query.decodeString()
	return response, err
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
