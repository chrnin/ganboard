package ganboard

import "encoding/json"

// CreateSubtask https://docs.kanboard.org/en/latest/api/subtask_procedures.html#createsubtask
func (c *Client) CreateSubtask(params SubtaskParams) (int, error) {
	query := request{
		Client: c,
		Method: "createSubtask",
		Params: params,
	}
	response, err := query.decodeInt()
	return response, err
}

// GetSubtask https://docs.kanboard.org/en/latest/api/subtask_procedures.html#getsubtask
func (c *Client) GetSubtask(subtaskID int) (Subtask, error) {
	query := request{
		Client: c,
		Method: "getSubtask",
		Params: map[string]int{
			"subtask_id": subtaskID,
		},
	}
	response, err := query.decodeSubtask()
	return response, err
}

// GetAllSubtasks https://docs.kanboard.org/en/latest/api/subtask_procedures.html#getallsubtasks
func (c *Client) GetAllSubtasks(taskID int) ([]Subtask, error) {
	query := request{
		Client: c,
		Method: "getAllSubtasks",
		Params: map[string]int{
			"task_id": taskID,
		},
	}
	response, err := query.decodeSubtasks()
	return response, err
}

// UpdateSubtask https://docs.kanboard.org/en/latest/api/subtask_procedures.html#updatesubtask
func (c *Client) UpdateSubtask(params SubtaskParams) (bool, error) {
	query := request{
		Client: c,
		Method: "updateSubtask",
		Params: params,
	}
	response, err := query.decodeBoolean()
	return response, err
}

// RemoveSubtask https://docs.kanboard.org/en/latest/api/subtask_procedures.html#removesubtask
func (c *Client) RemoveSubtask(subtaskID int) (bool, error) {
	query := request{
		Client: c,
		Method: "removeSubtask",
		Params: map[string]int{
			"subtask_id": subtaskID,
		},
	}
	response, err := query.decodeBoolean()
	return response, err
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

func (r *request) decodeSubtasks() ([]Subtask, error) {
	rsp, err := r.Client.Request(*r)
	if err != nil {
		return nil, err
	}

	body := struct {
		JSONRPC string    `json:"jsonrpc"`
		ID      int       `json:"id"`
		Result  []Subtask `json:"result"`
	}{}

	err = json.NewDecoder(rsp.Body).Decode(&body)
	return body.Result, err
}

func (r *request) decodeSubtask() (Subtask, error) {
	rsp, err := r.Client.Request(*r)
	if err != nil {
		return Subtask{}, err
	}

	body := struct {
		JSONRPC string  `json:"jsonrpc"`
		ID      int     `json:"id"`
		Result  Subtask `json:"result"`
	}{}

	err = json.NewDecoder(rsp.Body).Decode(&body)
	return body.Result, err
}
