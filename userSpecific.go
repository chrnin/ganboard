package ganboard

import (
	"encoding/json"
)

// GetMe https://docs.kanboard.org/en/latest/api/me_procedures.html#getme
func (c *Client) GetMe() (User, error) {
	request := request{
		JSONRPC: "2.0",
		Method:  "getMe",
		ID:      1,
	}

	rsp, err := c.Request(request)
	body := responseUser{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// GetMyDashboard https://docs.kanboard.org/en/latest/api/me_procedures.html#getmydashboard
// FIXME documentation doesn't fit result.
func (c *Client) GetMyDashboard() (interface{}, error) {
	request := request{
		JSONRPC: "2.0",
		Method:  "getMyDashboard",
		ID:      1,
	}

	rsp, err := c.Request(request)
	body := responseInt{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body, err
}

// GetMyActivityStream https://docs.kanboard.org/en/latest/api/me_procedures.html#getmyactivitystream
func (c *Client) GetMyActivityStream() ([]Activity, error) {
	request := request{
		JSONRPC: "2.0",
		Method:  "getMyActivityStream",
		ID:      1,
	}

	rsp, err := c.Request(request)
	body := responseActivityStream{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// CreateMyPrivateProject https://docs.kanboard.org/en/latest/api/me_procedures.html#createmyprivateproject
func (c *Client) CreateMyPrivateProject(params PrivateProjectParams) (int, error) {
	request := request{
		JSONRPC: "2.0",
		Method:  "createMyPrivateProject",
		ID:      1,
		Params:  params,
	}

	rsp, err := c.Request(request)
	body := response{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result.(int), err
}

// GetMyProjectList https://docs.kanboard.org/en/latest/api/me_procedures.html#getmyprojectslist
func (c *Client) GetMyProjectList() (map[int]string, error) {
	request := request{
		JSONRPC: "2.0",
		Method:  "getMyProjectsList",
		ID:      1,
	}

	rsp, err := c.Request(request)
	body := responseProjectList{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// GetMyOverDueTasks https://docs.kanboard.org/en/latest/api/me_procedures.html#getmyoverduetasks
func (c *Client) GetMyOverDueTasks() ([]Task, error) {
	request := request{
		JSONRPC: "2.0",
		Method:  "getMyOverDueTasks",
		ID:      1,
	}

	rsp, err := c.Request(request)
	body := responseOverdueTasks{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// GetMyProjects https://docs.kanboard.org/en/latest/api/me_procedures.html#getmyprojects
func (c *Client) GetMyProjects() ([]Project, error) {
	request := request{
		JSONRPC: "2.0",
		Method:  "getMyProjects",
		ID:      1,
	}

	rsp, err := c.Request(request)
	body := responseProjects{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// PrivateProjectParams parameters for CreateMyPrivateProject
type PrivateProjectParams struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

type responseOverdueTasks struct {
	JSONRPC string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  []Task `json:"result"`
}

type responseProjectList struct {
	JSONRPC string         `json:"jsonrpc"`
	ID      int            `json:"id"`
	Result  map[int]string `json:"result"`
}

type responseDashboard struct {
	JSONRPC string    `json:"jsonrpc"`
	ID      int       `json:"id"`
	Result  Dashboard `json:"result"`
}

type responseActivityStream struct {
	JSONRPC string     `json:"jsonrpc"`
	ID      int        `json:"id"`
	Result  []Activity `json:"result"`
}

// Dashboard type
type Dashboard struct {
	Projects []Project   `json:"projects"`
	Tasks    []Task      `json:"tasks"`
	SubTasks interface{} `json:"subtasks"`
}

// Activity type
// FIXME no information on Changes structure
type Activity struct {
	ID             int           `json:"id,string"`
	DateCreation   int           `json:"date_creation,string"`
	EventName      string        `json:"event_name"`
	CreatorID      int           `json:"creator_id,string"`
	ProjectID      int           `json:"project_id,string"`
	TaskID         int           `json:"task_id,string"`
	AuthorUsername string        `json:"author_username"`
	AuthorName     string        `json:"author_name"`
	Email          string        `json:"email"`
	Task           Task          `json:"task"`
	Changes        []interface{} `json:"changes"`
	Author         string        `json:"author"`
	EventTitle     string        `json:"event_title"`
	EventContent   string        `json:"event_content"`
}
