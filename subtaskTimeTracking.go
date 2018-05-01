package ganboard

// HasSubtaskTimer https://docs.kanboard.org/en/latest/api/subtask_time_tracking_procedures.html#hassubtasktimer
func (c *Client) HasSubtaskTimer(subtaskID int, userID int) (bool, error) {
	query := request{
		Client: c,
		Method: "hasSubtaskTimer",
		Params: map[string]int{
			"subtask_id": subtaskID,
			"user_id":    userID,
		},
	}
	response, err := query.decodeBoolean()
	return response, err
}

// SetSubtaskStartTime https://docs.kanboard.org/en/latest/api/subtask_time_tracking_procedures.html#setsubtaskstarttime
func (c *Client) SetSubtaskStartTime(subtaskID int, userID int) (bool, error) {
	query := request{
		Client: c,
		Method: "setSubtaskStartTime",
		Params: map[string]int{
			"subtask_id": subtaskID,
			"user_id":    userID,
		},
	}
	response, err := query.decodeBoolean()
	return response, err
}

// SetSubtaskEndTime https://docs.kanboard.org/en/latest/api/subtask_time_tracking_procedures.html#setsubtaskendtime
func (c *Client) SetSubtaskEndTime(subtaskID int, userID int) (bool, error) {
	query := request{
		Client: c,
		Method: "setSubtaskEndTime",
		Params: map[string]int{
			"subtask_id": subtaskID,
			"user_id":    userID,
		},
	}
	response, err := query.decodeBoolean()
	return response, err
}

// GetSubtaskTimeSpent https://docs.kanboard.org/en/latest/api/subtask_time_tracking_procedures.html#getsubtasktimespent
func (c *Client) GetSubtaskTimeSpent(subtaskID int, userID int) (float64, error) {
	query := request{
		Client: c,
		Method: "getSubtaskTimeSpent",
		Params: map[string]int{
			"subtask_id": subtaskID,
			"user_id":    userID,
		},
	}
	response, err := query.decodeFloat64()
	return response, err
}
