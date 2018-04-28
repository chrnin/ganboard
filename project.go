package ganboard

import (
	"bytes"
	"encoding/json"
)

// GetAllProjects https://docs.kanboard.org/en/latest/api/project_procedures.html#getallprojects
func (c *Client) GetAllProjects() ([]Project, error) {
	b := new(bytes.Buffer)
	r := Request{
		JSONRPC: "2.0",
		Method:  "getAllProjects",
		ID:      "1",
	}
	json.NewEncoder(b).Encode(r)
	rsp, err := c.Request(b)

	body := ResponseAllProjects{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// GetProjectById https://docs.kanboard.org/en/latest/api/project_procedures.html#getprojectbyid
func (c *Client) GetProjectById(id int) (Project, error) {
	b := new(bytes.Buffer)
	r := Request{
		JSONRPC: "2.0",
		Method:  "getProjectById",
		ID:      "1",
	}
	r.Params.ProjectID = id

	json.NewEncoder(b).Encode(r)
	rsp, err := c.Request(b)

	body := ResponseProject{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// GetProjectByName https://docs.kanboard.org/en/latest/api/project_procedures.html#getprojectbyname
func (c *Client) GetProjectByName(name string) (Project, error) {
	b := new(bytes.Buffer)
	r := Request{
		JSONRPC: "2.0",
		Method:  "getProjectByName",
		ID:      "1",
	}
	r.Params.Name = name

	json.NewEncoder(b).Encode(r)
	rsp, err := c.Request(b)

	body := ResponseProject{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// GetProjectByIdentifier https://docs.kanboard.org/en/latest/api/project_procedures.html#getprojectbyname
func (c *Client) GetProjectByIdentifier(identifier string) (Project, error) {
	b := new(bytes.Buffer)
	r := Request{
		JSONRPC: "2.0",
		Method:  "getProjectByIdentifier",
		ID:      "1",
	}
	r.Params.Identifier = identifier

	json.NewEncoder(b).Encode(r)
	rsp, err := c.Request(b)

	body := ResponseProject{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// Request jsonrpc with params
type Request struct {
	JSONRPC string `json:"jsonrpc"`
	Method  string `json:"method"`
	ID      string `json:"id"`
	Params  struct {
		ProjectID  int    `json:"project_id,omitempty"`
		Name       string `json:"name,omitempty"`
		Identifier string `json:"identifier,omitempty"`
		Email      string `json:"email,omitempty"`
	} `json:"params,omitempty"`
}

// ResponseAllProjects jsonrpc response for slice of projects
type ResponseAllProjects struct {
	JSONRPC string    `json:"jsonrpc"`
	ID      int       `json:"id"`
	Result  []Project `json:"result"`
}

// ResponseProject jsonrpc response for project
type ResponseProject struct {
	JSONRPC string  `json:"jsonrpc"`
	ID      int     `json:"id"`
	Result  Project `json:"result"`
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
	URL                 struct {
		Board    string `json:"board"`
		Calendar string `json:"calendar"`
		List     string `json:"list"`
	} `json:"url"`
}
