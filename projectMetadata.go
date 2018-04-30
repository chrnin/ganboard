package ganboard

import "encoding/json"

// GetProjectMetadata https://docs.kanboard.org/en/latest/api/project_metadata_procedures.html#getprojectmetadata
func (c *Client) GetProjectMetadata(projectID int) (map[string]string, error) {
	request := request{
		JSONRPC: "2.0",
		Method:  "getProjectMetadata",
		ID:      1,
		Params: map[string]int{
			"project_id": projectID,
		},
	}

	rsp, err := c.Request(request)
	body := responseMapStringString{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// GetProjectMetadataByName https://docs.kanboard.org/en/latest/api/project_metadata_procedures.html#getprojectmetadatabyname
func (c *Client) GetProjectMetadataByName(projectID int, name string) (string, error) {
	request := request{
		JSONRPC: "2.0",
		Method:  "getProjectMetadataByName",
		ID:      1,
		Params: struct {
			ProjectID int    `json:"project_id"`
			Name      string `json:"name"`
		}{
			ProjectID: projectID,
			Name:      name,
		},
	}

	rsp, err := c.Request(request)
	body := responseString{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// SaveProjectMetadata https://docs.kanboard.org/en/latest/api/project_metadata_procedures.html#saveprojectmetadata
func (c *Client) SaveProjectMetadata(projectID int, values map[string]string) (bool, error) {
	request := request{
		JSONRPC: "2.0",
		Method:  "saveProjectMetadata",
		ID:      1,
		Params: struct {
			ProjectID int               `json:"project_id"`
			Values    map[string]string `json:"values"`
		}{
			ProjectID: projectID,
			Values:    values,
		},
	}

	rsp, err := c.Request(request)
	body := responseBoolean{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// RemoveProjectMetadata https://docs.kanboard.org/en/latest/api/project_metadata_procedures.html#removeprojectmetadata
func (c *Client) RemoveProjectMetadata(projectID int, name string) (bool, error) {
	request := request{
		JSONRPC: "2.0",
		Method:  "removeProjectMetadata",
		ID:      1,
		Params: struct {
			ProjectID int    `json:"project_id"`
			Name      string `json:"name"`
		}{
			ProjectID: projectID,
			Name:      name,
		},
	}

	rsp, err := c.Request(request)
	body := responseBoolean{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}
