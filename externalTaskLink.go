package ganboard

import "encoding/json"

// GetExternalTaskLinkTypes https://docs.kanboard.org/en/latest/api/external_task_link_procedures.html#getexternaltasklinktypes
func (c *Client) GetExternalTaskLinkTypes() (map[string]string, error) {
	query := request{
		Client: c,
		Method: "getExternalTaskLinkTypes",
	}
	response, err := query.decodeMapStringString()
	return response, err
}

// GetExternalTaskLinkProviderDependencies https://docs.kanboard.org/en/latest/api/external_task_link_procedures.html#getexternaltasklinkproviderdependencies
func (c *Client) GetExternalTaskLinkProviderDependencies(providerName string) (map[string]string, error) {
	query := request{
		Client: c,
		Method: "getExternalTaskLinkProviderDependencies",
		Params: map[string]string{
			"providerName": providerName,
		},
	}
	response, err := query.decodeMapStringString()
	return response, err
}

// CreateExternalTaskLink https://docs.kanboard.org/en/latest/api/external_task_link_procedures.html#createexternaltasklink
func (c *Client) CreateExternalTaskLink(taskID int, url string, dependency string, taskType string, title string) (int, error) {
	query := request{
		Client: c,
		Method: "createExternalTaskLink",
		Params: struct {
			TaskID     int    `json:"task_id,string"`
			URL        string `json:"url"`
			Dependency string `json:"dependency"`
			TypeTask   string `json:"type"`
			Title      string `json:"title"`
		}{
			TaskID:     taskID,
			URL:        url,
			Dependency: dependency,
			TypeTask:   taskType,
			Title:      title,
		},
	}
	response, err := query.decodeInt()
	return response, err
}

// UpdateExternalTaskLink https://docs.kanboard.org/en/latest/api/external_task_link_procedures.html#updateexternaltasklink
func (c *Client) UpdateExternalTaskLink(taskID int, linkID int, title string, url string, dependency string) (bool, error) {
	query := request{
		Client: c,
		Method: "updateExternalTaskLink",
		Params: struct {
			TaskID     int    `json:"task_id,string"`
			LinkID     int    `json:"link_id,string"`
			URL        string `json:"url"`
			Dependency string `json:"dependency"`
			Title      string `json:"title"`
		}{
			TaskID:     taskID,
			URL:        url,
			Dependency: dependency,
			Title:      title,
			LinkID:     linkID,
		},
	}
	response, err := query.decodeBoolean()
	return response, err
}

// GetExternalTaskLinkByID https://docs.kanboard.org/en/latest/api/external_task_link_procedures.html#getexternaltasklinkbyid
func (c *Client) GetExternalTaskLinkByID(taskID int, linkID int) (ExternalLink, error) {
	query := request{
		Client: c,
		Method: "getExternalTaskLinkById",
		Params: map[string]int{
			"task_id": taskID,
			"link_id": linkID,
		},
	}
	response, err := query.decodeExternalLink()
	return response, err
}

// GetAllExternalTaskLinks https://docs.kanboard.org/en/latest/api/external_task_link_procedures.html#getallexternaltasklinks
func (c *Client) GetAllExternalTaskLinks(taskID int) ([]ExternalLink, error) {
	query := request{
		Client: c,
		Method: "getAllExternalTaskLinks",
		Params: map[string]int{
			"task_id": taskID,
		},
	}
	response, err := query.decodeExternalLinks()
	return response, err
}

// RemoveExternalTaskLink https://docs.kanboard.org/en/latest/api/external_task_link_procedures.html#removeexternaltasklink
func (c *Client) RemoveExternalTaskLink(taskID int, linkID int) (bool, error) {
	query := request{
		Client: c,
		Method: "removeExternalTaskLink",
		Params: map[string]int{
			"task_id": taskID,
			"link_id": linkID,
		},
	}
	response, err := query.decodeBoolean()
	return response, err
}

// ExternalLink type for external tasks links
type ExternalLink struct {
	ID               int    `json:"id,string"`
	LinkType         string `json:"link_type"`
	Dependency       string `json:"dependency"`
	Title            string `json:"title"`
	URL              string `json:"url"`
	DateCreation     int    `json:"date_creation,string"`
	DateModification int    `json:"date_modification,string"`
	TaskID           int    `json:"task_id,string"`
	CreatorID        int    `json:"create_id,string"`
}

func (r *request) decodeExternalLink() (ExternalLink, error) {
	rsp, err := r.Client.Request(*r)

	if err != nil {
		return ExternalLink{}, err
	}

	body := struct {
		JSONRPC string       `json:"jsonrpc"`
		ID      int          `json:"id"`
		Result  ExternalLink `json:"result"`
	}{}

	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

func (r *request) decodeExternalLinks() ([]ExternalLink, error) {
	rsp, err := r.Client.Request(*r)

	if err != nil {
		return nil, err
	}

	body := struct {
		JSONRPC string         `json:"jsonrpc"`
		ID      int            `json:"id"`
		Result  []ExternalLink `json:"result"`
	}{}

	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}
