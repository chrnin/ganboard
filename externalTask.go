package ganboard

import "encoding/json"

// GetExternalTaskLinkTypes https://docs.kanboard.org/en/latest/api/external_task_link_procedures.html#getexternaltasklinktypes
func (c *Client) GetExternalTaskLinkTypes() (map[string]string, error) {
	r := request{
		JSONRPC: "2.0",
		Method:  "getExternalTaskLinkTypes",
		ID:      1,
	}

	rsp, err := c.Request(r)
	body := responseMapStringString{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// GetExternalTaskLinkProviderDependencies https://docs.kanboard.org/en/latest/api/external_task_link_procedures.html#getexternaltasklinkproviderdependencies
func (c *Client) GetExternalTaskLinkProviderDependencies(providerName string) (map[string]string, error) {
	r := request{
		JSONRPC: "2.0",
		Method:  "getExternalTaskLinkProviderDependencies",
		ID:      1,
		Params: map[string]string{
			"providerName": providerName,
		},
	}

	rsp, err := c.Request(r)
	body := responseMapStringString{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// CreateExternalTaskLink https://docs.kanboard.org/en/latest/api/external_task_link_procedures.html#createexternaltasklink
func (c *Client) CreateExternalTaskLink(taskID int, url string, dependency string, typeTask string, title string) (int, error) {
	r := request{
		JSONRPC: "2.0",
		Method:  "createExternalTaskLink",
		ID:      1,
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
			TypeTask:   typeTask,
			Title:      title,
		},
	}

	rsp, err := c.Request(r)
	body := responseInt{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// UpdateExternalTaskLink https://docs.kanboard.org/en/latest/api/external_task_link_procedures.html#updateexternaltasklink
func (c *Client) UpdateExternalTaskLink(taskID int, linkID int, title string, url string, dependency string) (bool, error) {
	r := request{
		JSONRPC: "2.0",
		Method:  "updateExternalTaskLink",
		ID:      1,
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

	rsp, err := c.Request(r)
	body := responseBoolean{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// GetExternalTaskLinkByID https://docs.kanboard.org/en/latest/api/external_task_link_procedures.html#getexternaltasklinkbyid
func (c *Client) GetExternalTaskLinkByID(taskID int, linkID int) (Link, error) {
	r := request{
		JSONRPC: "2.0",
		Method:  "getExternalTaskLinkById",
		ID:      1,
		Params: map[string]int{
			"task_id": taskID,
			"link_id": linkID,
		},
	}

	rsp, err := c.Request(r)
	body := responseLink{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// GetAllExternalTaskLinks https://docs.kanboard.org/en/latest/api/external_task_link_procedures.html#getallexternaltasklinks
func (c *Client) GetAllExternalTaskLinks(taskID int) ([]Link, error) {
	r := request{
		JSONRPC: "2.0",
		Method:  "getAllExternalTaskLinks",
		ID:      1,
		Params: map[string]int{
			"task_id": taskID,
		},
	}

	rsp, err := c.Request(r)
	body := responseLinks{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// RemoveExternalTaskLink https://docs.kanboard.org/en/latest/api/external_task_link_procedures.html#removeexternaltasklink
func (c *Client) RemoveExternalTaskLink(taskID int, linkID int) (bool, error) {
	r := request{
		JSONRPC: "2.0",
		Method:  "removeExternalTaskLink",
		ID:      1,
		Params: map[string]int{
			"task_id": taskID,
			"link_id": linkID,
		},
	}

	rsp, err := c.Request(r)
	body := responseBoolean{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// Link for external tasks type
type Link struct {
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

type responseLink struct {
	JSONRPC string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  Link   `json:"result"`
}

type responseLinks struct {
	JSONRPC string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  []Link `json:"result"`
}
