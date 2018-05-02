package ganboard

import "encoding/json"

// GetAllTags https://docs.kanboard.org/en/latest/api/tags_procedures.html#getalltags
func (c *Client) GetAllTags() ([]Tag, error) {
	query := request{
		Client: c,
		Method: "getAllTags",
	}
	response, err := query.decodeTags()
	return response, err
}

// GetTagsByProject https://docs.kanboard.org/en/latest/api/tags_procedures.html#gettagsbyproject
func (c *Client) GetTagsByProject(projectID int) ([]Tag, error) {
	query := request{
		Client: c,
		Method: "getTagsByProject",
		Params: map[string]int{
			"project_id": projectID,
		},
	}
	response, err := query.decodeTags()
	return response, err
}

// CreateTag https://docs.kanboard.org/en/latest/api/tags_procedures.html#createtag
func (c *Client) CreateTag(projectID int, tag string) (int, error) {
	query := request{
		Client: c,
		Method: "createTag",
		Params: struct {
			ProjectID int    `json:"project_id,string"`
			Tag       string `json:"tag"`
		}{
			ProjectID: projectID,
			Tag:       tag,
		},
	}
	response, err := query.decodeInt()
	return response, err
}

// UpdateTag https://docs.kanboard.org/en/latest/api/tags_procedures.html#updatetag
func (c *Client) UpdateTag(projectID int, tag string) (bool, error) {
	query := request{
		Client: c,
		Method: "updateTag",
		Params: struct {
			ProjectID int    `json:"project_id,string"`
			Tag       string `json:"tag"`
		}{
			ProjectID: projectID,
			Tag:       tag,
		},
	}
	response, err := query.decodeBoolean()
	return response, err
}

// RemoveTag https://docs.kanboard.org/en/latest/api/tags_procedures.html#removetag
func (c *Client) RemoveTag(tagID int) (bool, error) {
	query := request{
		Client: c,
		Method: "removeTag",
		Params: map[string]int{
			"tag_id": tagID,
		},
	}
	response, err := query.decodeBoolean()
	return response, err
}

// setTaskTags https://docs.kanboard.org/en/latest/api/tags_procedures.html#settasktags
func (c *Client) setTaskTags(projectID int, taskID int, tags []string) (bool, error) {
	query := request{
		Client: c,
		Method: "setTaskTags",
		Params: struct {
			ProjectID int      `json:"project_id,string"`
			TaskID    int      `json:"task_id"`
			Tags      []string `json:"tags"`
		}{
			ProjectID: projectID,
			TaskID:    taskID,
			Tags:      tags,
		},
	}
	response, err := query.decodeBoolean()
	return response, err
}

// GetTaskTags https://docs.kanboard.org/en/latest/api/tags_procedures.html#gettasktags
func (c *Client) GetTaskTags(taskID int) (map[int]string, error) {
	query := request{
		Client: c,
		Method: "getTaskTags",
		Params: map[string]int{
			"task_id": taskID,
		},
	}
	response, err := query.decodeMapIntString()
	return response, err
}

// Tag type
type Tag struct {
	ID        int    `json:"id,string"`
	Name      string `json:"string"`
	ProjectID int    `json:"project_id,string"`
}

func (r *request) decodeTag() (Tag, error) {
	rsp, err := r.Client.Request(*r)
	if err != nil {
		return Tag{}, err
	}

	body := struct {
		JSONRPC string `json:"jsonrpc"`
		ID      int    `json:"id"`
		Result  Tag    `json:"result"`
	}{}

	err = json.NewDecoder(rsp.Body).Decode(&body)
	return body.Result, err
}

func (r *request) decodeTags() ([]Tag, error) {
	rsp, err := r.Client.Request(*r)
	if err != nil {
		return nil, err
	}

	body := struct {
		JSONRPC string `json:"jsonrpc"`
		ID      int    `json:"id"`
		Result  []Tag  `json:"result"`
	}{}

	err = json.NewDecoder(rsp.Body).Decode(&body)
	return body.Result, err
}
