package ganboard

// GetTaskMetadata https://docs.kanboard.org/en/latest/api/task_metadata_procedures.html#gettaskmetadata
func (c *Client) GetTaskMetadata(taskID int) (map[string]string, error) {
	query := request{
		Client: c,
		Method: "getTaskMetadata",
		Params: map[string]int{
			"task_id": taskID,
		},
	}
	response, err := query.decodeMapStringString()
	return response, err
}

// GetTaskMetadataByName https://docs.kanboard.org/en/latest/api/task_metadata_procedures.html#gettaskmetadatabyname
func (c *Client) GetTaskMetadataByName(taskID int, name string) (string, error) {
	query := request{
		Client: c,
		Method: "getTaskMetadataByName",
		Params: struct {
			TaskID int    `json:"task_id"`
			Name   string `json:"name"`
		}{
			TaskID: taskID,
			Name:   name,
		},
	}
	response, err := query.decodeString()
	return response, err
}

// SaveTaskMetadata https://docs.kanboard.org/en/latest/api/task_metadata_procedures.html#savetaskmetadata
// FIXME metadata model not clear in documentation
func (c *Client) SaveTaskMetadata(taskID int, name string, value string) (bool, error) {
	query := request{
		Client: c,
		Method: "saveTaskMetadata",
		Params: struct {
			TaskID   int               `json:"task_id"`
			Metadata map[string]string `json:"metadata"`
		}{
			TaskID: taskID,
			Metadata: map[string]string{
				name: value,
			},
		},
	}
	response, err := query.decodeBoolean()
	return response, err
}

// RemoveTaskMetadata https://docs.kanboard.org/en/latest/api/task_metadata_procedures.html#removetaskmetadata
func (c *Client) RemoveTaskMetadata(taskID int, name string) (bool, error) {
	query := request{
		Client: c,
		Method: "removeTaskMetadata",
		Params: struct {
			TaskID int    `json:"task_id"`
			Name   string `json:"name"`
		}{
			TaskID: taskID,
			Name:   name,
		},
	}
	response, err := query.decodeBoolean()
	return response, err
}
