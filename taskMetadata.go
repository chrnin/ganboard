package ganboard

import "encoding/json"

// GetTaskMetadata https://docs.kanboard.org/en/latest/api/task_metadata_procedures.html#gettaskmetadata
func (c *Client) GetTaskMetadata(taskID int) (map[string]string, error) {
	r := request{
		JSONRPC: "2.0",
		Method:  "getTaskMetadata",
		ID:      1,
		Params: map[string]int{
			"task_id": taskID,
		},
	}

	rsp, err := c.Request(r)
	body := responseMapStringString{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// GetTaskMetadataByName https://docs.kanboard.org/en/latest/api/task_metadata_procedures.html#gettaskmetadatabyname
func (c *Client) GetTaskMetadataByName(taskID int, name string) (string, error) {
	r := request{
		JSONRPC: "2.0",
		Method:  "getTaskMetadataByName",
		ID:      1,
		Params: struct {
			TaskID int    `json:"task_id"`
			Name   string `json:"name"`
		}{
			TaskID: taskID,
			Name:   name,
		},
	}

	rsp, err := c.Request(r)
	body := responseString{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// SaveTaskMetadata https://docs.kanboard.org/en/latest/api/task_metadata_procedures.html#savetaskmetadata
// FIXME metadata model not clear in documentation
func (c *Client) SaveTaskMetadata(taskID int, name string, value string) (bool, error) {
	r := request{
		JSONRPC: "2.0",
		Method:  "saveTaskMetadata",
		ID:      1,
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

	rsp, err := c.Request(r)
	body := responseBoolean{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// RemoveTaskMetadata https://docs.kanboard.org/en/latest/api/task_metadata_procedures.html#removetaskmetadata
func (c *Client) RemoveTaskMetadata(taskID int, name string) (bool, error) {
	r := request{
		JSONRPC: "2.0",
		Method:  "RemoveTaskMetadata",
		ID:      1,
		Params: struct {
			TaskID int    `json:"task_id"`
			Name   string `json:"name"`
		}{
			TaskID: taskID,
			Name:   name,
		},
	}

	rsp, err := c.Request(r)
	body := responseBoolean{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}
