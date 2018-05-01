package ganboard

import (
	"encoding/json"
)

// GetMe https://docs.kanboard.org/en/latest/api/me_procedures.html#getme
func (c *Client) GetMe() (User, error) {
	query := request{
		Client: c,
		Method: "getMe",
	}
	response, err := query.decodeUser()
	return response, err
}

// GetMyDashboard https://docs.kanboard.org/en/latest/api/me_procedures.html#getmydashboard
// FIXME documentation doesn't fit result.
func (c *Client) GetMyDashboard() (interface{}, error) {
	query := request{
		Client: c,
		Method: "getMyDashboard",
		ID:     1,
	}
	response, err := query.decodeInterface()
	return response, err
}

// GetMyActivityStream https://docs.kanboard.org/en/latest/api/me_procedures.html#getmyactivitystream
func (c *Client) GetMyActivityStream() ([]Activity, error) {
	query := request{
		Client: c,
		Method: "getMyActivityStream",
	}
	response, err := query.decodeActivities()
	return response, err
}

// CreateMyPrivateProject https://docs.kanboard.org/en/latest/api/me_procedures.html#createmyprivateproject
func (c *Client) CreateMyPrivateProject(params PrivateProjectParams) (int, error) {
	query := request{
		Client: c,
		Method: "createMyPrivateProject",
		Params: params,
	}
	response, err := query.decodeInt()
	return response, err
}

// GetMyProjectList https://docs.kanboard.org/en/latest/api/me_procedures.html#getmyprojectslist
func (c *Client) GetMyProjectList() (map[int]string, error) {
	query := request{
		Client: c,
		Method: "getMyProjectsList",
	}
	response, err := query.decodeMapIntString()
	return response, err
}

// GetMyOverDueTasks https://docs.kanboard.org/en/latest/api/me_procedures.html#getmyoverduetasks
func (c *Client) GetMyOverDueTasks() ([]Task, error) {
	query := request{
		Client: c,
		Method: "getMyOverDueTasks",
	}
	response, err := query.decodeTasks()
	return response, err
}

// GetMyProjects https://docs.kanboard.org/en/latest/api/me_procedures.html#getmyprojects
func (c *Client) GetMyProjects() ([]Project, error) {
	query := request{
		Client: c,
		Method: "getMyProjects",
	}
	response, err := query.decodeProjects()
	return response, err
}

// PrivateProjectParams parameters for CreateMyPrivateProject
type PrivateProjectParams struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
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

func (r *request) decodeDashboard() (Dashboard, error) {
	rsp, err := r.Client.Request(*r)
	if err != nil {
		return Dashboard{}, err
	}

	body := struct {
		JSONRPC string    `json:"jsonrpc"`
		ID      int       `json:"id"`
		Result  Dashboard `json:"result"`
	}{}

	err = json.NewDecoder(rsp.Body).Decode(&body)
	return body.Result, err
}

func (r *request) decodeActivities() ([]Activity, error) {
	rsp, err := r.Client.Request(*r)
	if err != nil {
		return nil, err
	}

	body := struct {
		JSONRPC string     `json:"jsonrpc"`
		ID      int        `json:"id"`
		Result  []Activity `json:"result"`
	}{}

	err = json.NewDecoder(rsp.Body).Decode(&body)
	return body.Result, err
}
