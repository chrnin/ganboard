package ganboard

import (
	"encoding/json"
)

// CreateProject https://docs.kanboard.org/en/latest/api/project_procedures.html#createprojects
func (c *Client) CreateProject(params ProjectParams) (int, error) {
	request := request{
		JSONRPC: "2.0",
		Method:  "createProject",
		ID:      "1",
		Params:  params,
	}

	rsp, err := c.Request(request)
	body := projectCreateResponse{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// GetAllProjects https://docs.kanboard.org/en/latest/api/project_procedures.html#getallprojects
func (c *Client) GetAllProjects() ([]Project, error) {
	request := request{
		JSONRPC: "2.0",
		Method:  "getAllProjects",
		ID:      "1",
	}

	rsp, err := c.Request(request)
	body := projectsResponse{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// GetProjectByID https://docs.kanboard.org/en/latest/api/project_procedures.html#getprojectbyid
func (c *Client) GetProjectByID(id int) (Project, error) {
	r := request{
		JSONRPC: "2.0",
		Method:  "getProjectById",
		ID:      "1",
		Params: ProjectParams{
			ProjectID: id,
		},
	}

	rsp, err := c.Request(r)
	body := projectResponse{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// GetProjectByName https://docs.kanboard.org/en/latest/api/project_procedures.html#getprojectbyname
func (c *Client) GetProjectByName(name string) (Project, error) {
	r := request{
		JSONRPC: "2.0",
		Method:  "getProjectByName",
		ID:      "1",
		Params: ProjectParams{
			Name: name,
		},
	}

	rsp, err := c.Request(r)
	body := projectResponse{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// GetProjectByIdentifier https://docs.kanboard.org/en/latest/api/project_procedures.html#getprojectbyname
func (c *Client) GetProjectByIdentifier(identifier string) (Project, error) {
	r := request{
		JSONRPC: "2.0",
		Method:  "getProjectByIdentifier",
		ID:      "1",
		Params: ProjectParams{
			Identifier: identifier,
		},
	}

	rsp, err := c.Request(r)
	body := projectResponse{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// UpdateProject https://docs.kanboard.org/en/latest/api/project_procedures.html#updateproject
func (c *Client) UpdateProject(params ProjectParams) (bool, error) {
	r := request{
		JSONRPC: "2.0",
		Method:  "updateProject",
		ID:      "1",
		Params:  params,
	}

	rsp, err := c.Request(r)
	body := response{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	if body.Result == nil {
		return false, err
	}
	return true, err
}

// RemoveProject https://docs.kanboard.org/en/latest/api/project_procedures.html#removeproject
func (c *Client) RemoveProject(projectID int) (bool, error) {
	r := request{
		JSONRPC: "2.0",
		Method:  "updateProject",
		ID:      "1",
		Params: ProjectParams{
			ProjectID: projectID,
		},
	}

	rsp, err := c.Request(r)
	body := response{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	if body.Result == nil {
		return false, err
	}
	return true, err
}

// EnableProject https://docs.kanboard.org/en/latest/api/project_procedures.html#enableproject
func (c *Client) EnableProject(projectID int) (bool, error) {
	r := request{
		JSONRPC: "2.0",
		Method:  "enableProject",
		ID:      "1",
		Params: ProjectParams{
			ProjectID: projectID,
		},
	}

	rsp, err := c.Request(r)
	body := response{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	if body.Result == nil {
		return false, err
	}
	return true, err
}

// DisableProject https://docs.kanboard.org/en/latest/api/project_procedures.html#enableproject
func (c *Client) DisableProject(projectID int) (bool, error) {
	r := request{
		JSONRPC: "2.0",
		Method:  "disableProject",
		ID:      "1",
		Params: ProjectParams{
			ProjectID: projectID,
		},
	}

	rsp, err := c.Request(r)
	body := response{}
	err = json.NewDecoder(rsp.Body).Decode(&body)
	if body.Result == nil {
		return false, err
	}
	return true, err
}

// GetProjectActivity https://docs.kanboard.org/en/latest/api/project_procedures.html#getprojectactivity
func (c *Client) GetProjectActivity(projectID int) (interface{}, error) {
	r := request{
		JSONRPC: "2.0",
		Method:  "getProjectActivity",
		ID:      "1",
		Params: ProjectParams{
			ProjectID: projectID,
		},
	}

	rsp, err := c.Request(r)
	body := new(interface{})
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body, err
}

// GetProjectActivities https://docs.kanboard.org/en/latest/api/project_procedures.html#getprojectactivity
func (c *Client) GetProjectActivities(projectID []int) (interface{}, error) {
	r := request{
		JSONRPC: "2.0",
		Method:  "getProjectActivity",
		ID:      "1",
		Params: map[string][]int{
			"project_ids": projectID,
		},
	}

	rsp, err := c.Request(r)
	body := new(interface{})
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body, err
}

// TODO: create Event type to decode Json in usable structs. Missing data type documentation.

// ProjectParams parameters for UpdateProject()
type ProjectParams struct {
	ProjectID   int    `json:"project_id,omitempty"`
	Name        string `json:"name,omitempty"`
	Identifier  string `json:"identifier,omitempty"`
	Email       string `json:"email,omitempty"`
	Description string `json:"decription,omitempty"`
}

type projectsResponse struct {
	JSONRPC string    `json:"jsonrpc"`
	ID      int       `json:"id"`
	Result  []Project `json:"result"`
}

type projectResponse struct {
	JSONRPC string  `json:"jsonrpc"`
	ID      int     `json:"id"`
	Result  Project `json:"result"`
}

type projectCreateResponse struct {
	JSONRPC string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  int    `json:"result"`
}

// Project reflects getAllProjects method
type Project struct {
	ID                  int    `json:"id,string"`
	Name                string `json:"name"`
	IsActive            int    `json:"is_active,string"`
	Token               string `json:"token"`
	LastModified        int    `json:"last_modified,string"`
	IsPublic            int    `json:"is_public,string"`
	IsPrivate           int    `json:"is_private,string"`
	DefaultSwimlane     string `json:"default_swimlane"`
	ShowDefaultSwimlane int    `json:"show_default_swimlane,string"`
	Description         string `json:"description"`
	Identifier          string `json:"identifier"`
	Columns             []struct {
		ID          int    `json:"id,string"`
		Title       string `json:"title"`
		Position    int    `json:"position,string"`
		ProjectID   int    `json:"project_id,string"`
		TaskLimit   int    `json:"task_limit,string"`
		Description string `json:"description"`
	} `json:"columns"`
	URL struct {
		Board    string `json:"board"`
		Calendar string `json:"calendar"`
		List     string `json:"list"`
	} `json:"url"`
}
