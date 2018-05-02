package ganboard

import "encoding/json"

// CreateTaskLink https://docs.kanboard.org/en/latest/api/internal_task_link_procedures.html#createtasklink
func (c *Client) CreateTaskLink(taskID int, oppositeTaskID int, linkID int) (int, error) {
	query := request{
		Client: c,
		Method: "createTaskLink",
		Params: map[string]int{
			"task_id":          taskID,
			"opposite_task_id": oppositeTaskID,
			"link_id":          linkID,
		},
	}
	response, err := query.decodeInt()
	return response, err
}

// UpdateTaskLink https://docs.kanboard.org/en/latest/api/internal_task_link_procedures.html#updatetasklink
func (c *Client) UpdateTaskLink(taskLinkID int, taskID int, oppositeTaskID int, linkID int) (int, error) {
	query := request{
		Client: c,
		Method: "updateTaskLink",
		Params: map[string]int{
			"task_link_id":     taskLinkID,
			"task_id":          taskID,
			"opposite_task_id": oppositeTaskID,
			"link_id":          linkID,
		},
	}
	response, err := query.decodeInt()
	return response, err
}

// GetTaskLinkByID https://docs.kanboard.org/en/latest/api/internal_task_link_procedures.html#gettasklinkbyid
func (c *Client) GetTaskLinkByID(taskLinkID int) (InternalLink, error) {
	query := request{
		Client: c,
		Method: "getTaskLinkById",
		Params: map[string]int{
			"task_link_id": taskLinkID,
		},
	}
	response, err := query.decodeInternalLink()
	return response, err
}

// GetAllTaskLinks https://docs.kanboard.org/en/latest/api/internal_task_link_procedures.html#getalltasklinks
func (c *Client) GetAllTaskLinks() ([]InternalLink, error) {
	query := request{
		Client: c,
		Method: "getAllTaskLinks",
	}
	response, err := query.decodeInternalLinks()
	return response, err
}

// RemoveTaskLink https://docs.kanboard.org/en/latest/api/internal_task_link_procedures.html#removetasklink
func (c *Client) RemoveTaskLink(taskLinkID int) (bool, error) {
	query := request{
		Client: c,
		Method: "removeTaskLink",
		Params: map[string]int{
			"task_link_id": taskLinkID,
		},
	}
	response, err := query.decodeBoolean()
	return response, err
}

// InternalLink type
type InternalLink struct {
	ID             int `json:"id,string"`
	LinkID         int `json:"link_id,string"`
	TaskID         int `json:"task_id,string"`
	OppositeTaskID int `json:"opposite_task_id"`
}

func (r *request) decodeInternalLink() (InternalLink, error) {
	rsp, err := r.Client.Request(*r)
	if err != nil {
		return InternalLink{}, err
	}

	body := struct {
		JSONRPC string       `json:"jsonrpc"`
		ID      int          `json:"id"`
		Result  InternalLink `json:"result"`
	}{}

	err = json.NewDecoder(rsp.Body).Decode(&body)
	return body.Result, err
}

func (r *request) decodeInternalLinks() ([]InternalLink, error) {
	rsp, err := r.Client.Request(*r)
	if err != nil {
		return nil, err
	}

	body := struct {
		JSONRPC string         `json:"jsonrpc"`
		ID      int            `json:"id"`
		Result  []InternalLink `json:"result"`
	}{}

	err = json.NewDecoder(rsp.Body).Decode(&body)
	return body.Result, err
}
