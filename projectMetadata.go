package ganboard

// GetProjectMetadata https://docs.kanboard.org/en/latest/api/project_metadata_procedures.html#getprojectmetadata
func (c *Client) GetProjectMetadata(projectID int) (map[string]string, error) {
	query := request{
		Client: c,
		Method: "getProjectMetadata",
		Params: map[string]int{
			"project_id": projectID,
		},
	}
	response, err := query.decodeMapStringString()
	return response, err
}

// GetProjectMetadataByName https://docs.kanboard.org/en/latest/api/project_metadata_procedures.html#getprojectmetadatabyname
func (c *Client) GetProjectMetadataByName(projectID int, name string) (string, error) {
	query := request{
		Client: c,
		Method: "getProjectMetadataByName",
		ID:     1,
		Params: struct {
			ProjectID int    `json:"project_id"`
			Name      string `json:"name"`
		}{
			ProjectID: projectID,
			Name:      name,
		},
	}
	response, err := query.decodeString()
	return response, err
}

// SaveProjectMetadata https://docs.kanboard.org/en/latest/api/project_metadata_procedures.html#saveprojectmetadata
func (c *Client) SaveProjectMetadata(projectID int, values map[string]string) (bool, error) {
	query := request{
		Client: c,
		Method: "saveProjectMetadata",
		Params: struct {
			ProjectID int               `json:"project_id"`
			Values    map[string]string `json:"values"`
		}{
			ProjectID: projectID,
			Values:    values,
		},
	}
	response, err := query.decodeBoolean()
	return response, err
}

// RemoveProjectMetadata https://docs.kanboard.org/en/latest/api/project_metadata_procedures.html#removeprojectmetadata
func (c *Client) RemoveProjectMetadata(projectID int, name string) (bool, error) {
	query := request{
		Client: c,
		Method: "removeProjectMetadata",
		Params: struct {
			ProjectID int    `json:"project_id"`
			Name      string `json:"name"`
		}{
			ProjectID: projectID,
			Name:      name,
		},
	}
	response, err := query.decodeBoolean()
	return response, err
}
