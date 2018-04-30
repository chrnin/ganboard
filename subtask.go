package ganboard

import "encoding/json"

// CreateSubtask https://docs.kanboard.org/en/latest/api/subtask_procedures.html#createsubtask
func (c *Client) CreateSubtask(params SubtaskParams) (int, error) {
	r := request{
		JSONRPC: "2.0",
		Method:  "createSubtask",
		ID:      1,
		Params:  params,
	}

	rsp, err := c.Request(r)
	body := responseInt{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// GetSubtask https://docs.kanboard.org/en/latest/api/subtask_procedures.html#getsubtask
func (c *Client) GetSubtask(subtaskID int) (Subtask, error) {
	r := request{
		JSONRPC: "2.0",
		Method:  "getSubtask",
		ID:      1,
		Params: map[string]int{
			"subtask_id": subtaskID,
		},
	}

	rsp, err := c.Request(r)
	body := responseSubtask{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// GetAllSubtasks https://docs.kanboard.org/en/latest/api/subtask_procedures.html#getallsubtasks
func (c *Client) GetAllSubtasks(taskID int) ([]Subtask, error) {
	r := request{
		JSONRPC: "2.0",
		Method:  "getAllSubtasks",
		ID:      1,
		Params: map[string]int{
			"task_id": taskID,
		},
	}

	rsp, err := c.Request(r)
	body := responseSubtasks{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// UpdateSubtask https://docs.kanboard.org/en/latest/api/subtask_procedures.html#updatesubtask
func (c *Client) UpdateSubtask(params SubtaskParams) (bool, error) {
	r := request{
		JSONRPC: "2.0",
		Method:  "updateSubtask",
		ID:      1,
		Params:  params,
	}

	rsp, err := c.Request(r)
	body := responseBoolean{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// RemoveSubtask https://docs.kanboard.org/en/latest/api/subtask_procedures.html#removesubtask
func (c *Client) RemoveSubtask(subtaskID int) (bool, error) {
	r := request{
		JSONRPC: "2.0",
		Method:  "removeSubtask",
		ID:      1,
		Params: map[string]int{
			"subtask_id": subtaskID,
		},
	}

	rsp, err := c.Request(r)
	body := responseBoolean{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// SubtaskParams input for CreateSubtask
type SubtaskParams struct {
	ID            int    `json:"id,string,omitempty"`
	TaskID        int    `json:"task_id,string"`
	Title         string `json:"title"`
	UserID        int    `json:"user_id,string,omitempty"`
	TimeEstimated int    `json:"time_estimated,string,omitempty"`
	TimeSpent     int    `json:"time_spent,string,omitempty"`
	Status        int    `json:"status,string,omitempty"`
}

type responseSubtasks struct {
	JSONRPC string    `json:"jsonrpc"`
	ID      int       `json:"id"`
	Result  []Subtask `json:"result"`
}

type responseSubtask struct {
	JSONRPC string  `json:"jsonrpc"`
	ID      int     `json:"id"`
	Result  Subtask `json:"result"`
}

// Subtask type
type Subtask struct {
	ID            int    `json:"id,string"`
	TaskID        int    `json:"task_id,string"`
	Title         string `json:"title"`
	UserID        int    `json:"user_id,string,omitempty"`
	Username      string `json:"username"`
	Name          string `json:"name"`
	StatusName    string `json:"status_name"`
	TimeEstimated int    `json:"time_estimated,string,omitempty"`
	TimeSpent     int    `json:"time_spent,string,omitempty"`
	Status        int    `json:"status,string,omitempty"`
}
