package ganboard

import "encoding/json"

// HasSubtaskTimer https://docs.kanboard.org/en/latest/api/subtask_time_tracking_procedures.html#hassubtasktimer
func (c *Client) HasSubtaskTimer(subtaskID int, userID int) (bool, error) {
	r := request{
		JSONRPC: "2.0",
		Method:  "hasSubtaskTimer",
		ID:      1,
		Params: map[string]int{
			"subtask_id": subtaskID,
			"user_id":    userID,
		},
	}

	rsp, err := c.Request(r)
	body := responseBoolean{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// SetSubtaskStartTime https://docs.kanboard.org/en/latest/api/subtask_time_tracking_procedures.html#setsubtaskstarttime
func (c *Client) SetSubtaskStartTime(subtaskID int, userID int) (bool, error) {
	r := request{
		JSONRPC: "2.0",
		Method:  "setSubtaskStartTime",
		ID:      1,
		Params: map[string]int{
			"subtask_id": subtaskID,
			"user_id":    userID,
		},
	}

	rsp, err := c.Request(r)
	body := responseBoolean{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// SetSubtaskEndTime https://docs.kanboard.org/en/latest/api/subtask_time_tracking_procedures.html#setsubtaskendtime
func (c *Client) SetSubtaskEndTime(subtaskID int, userID int) (bool, error) {
	r := request{
		JSONRPC: "2.0",
		Method:  "setSubtaskEndTime",
		ID:      1,
		Params: map[string]int{
			"subtask_id": subtaskID,
			"user_id":    userID,
		},
	}

	rsp, err := c.Request(r)
	body := responseBoolean{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// GetSubtaskTimeSpent https://docs.kanboard.org/en/latest/api/subtask_time_tracking_procedures.html#getsubtasktimespent
func (c *Client) GetSubtaskTimeSpent(subtaskID int, userID int) (float64, error) {
	r := request{
		JSONRPC: "2.0",
		Method:  "getSubtaskTimeSpent",
		ID:      1,
		Params: map[string]int{
			"subtask_id": subtaskID,
			"user_id":    userID,
		},
	}

	rsp, err := c.Request(r)
	body := responseFloat64{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}
