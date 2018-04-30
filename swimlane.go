package ganboard

import "encoding/json"

// GetActiveSwimlanes https://docs.kanboard.org/en/latest/api/swimlane_procedures.html#swimlane-api-procedures
func (c *Client) GetActiveSwimlanes(projectID int) ([]Swimlane, error) {
	r := request{
		JSONRPC: "2.0",
		Method:  "getActiveSwimlanes",
		ID:      1,
		Params: map[string]int{
			"project_id": projectID,
		},
	}

	rsp, err := c.Request(r)
	body := responseSwimlanes{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// GetAllSwimlanes https://docs.kanboard.org/en/latest/api/swimlane_procedures.html#getallswimlanes
func (c *Client) GetAllSwimlanes(projectID int) ([]Swimlane, error) {
	r := request{
		JSONRPC: "2.0",
		Method:  "getAllSwimlanes",
		ID:      1,
		Params: map[string]int{
			"project_id": projectID,
		},
	}

	rsp, err := c.Request(r)
	body := responseSwimlanes{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// GetSwimlane https://docs.kanboard.org/en/latest/api/swimlane_procedures.html#getswimlane
func (c *Client) GetSwimlane(swimlaneID int) (Swimlane, error) {
	r := request{
		JSONRPC: "2.0",
		Method:  "getSwimlane",
		ID:      1,
		Params: map[string]int{
			"swimlane_id": swimlaneID,
		},
	}

	rsp, err := c.Request(r)
	body := responseSwimlane{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// GetSwimlaneByID https://docs.kanboard.org/en/latest/api/swimlane_procedures.html#getswimlanebyid
func (c *Client) GetSwimlaneByID(swimlaneID int) (Swimlane, error) {
	r := request{
		JSONRPC: "2.0",
		Method:  "getSwimlaneById",
		ID:      1,
		Params: map[string]int{
			"swimlane_id": swimlaneID,
		},
	}

	rsp, err := c.Request(r)
	body := responseSwimlane{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// GetSwimlaneByName https://docs.kanboard.org/en/latest/api/swimlane_procedures.html#getswimlanebyname
func (c *Client) GetSwimlaneByName(projectID int, name string) (Swimlane, error) {
	r := request{
		JSONRPC: "2.0",
		Method:  "getSwimlaneByName",
		ID:      1,
		Params: struct {
			ProjectID int    `json:"project_id"`
			Name      string `json:"Name"`
		}{
			ProjectID: projectID,
			Name:      name,
		},
	}

	rsp, err := c.Request(r)
	body := responseSwimlane{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// ChangeSwimlanePosition https://docs.kanboard.org/en/latest/api/swimlane_procedures.html#changeswimlaneposition
func (c *Client) ChangeSwimlanePosition(projectID int, swimlaneID int, position int) (bool, error) {
	r := request{
		JSONRPC: "2.0",
		Method:  "getSwimlaneByName",
		ID:      1,
		Params: map[string]int{
			"project_id":  projectID,
			"swimlane_id": swimlaneID,
			"position":    position,
		},
	}

	rsp, err := c.Request(r)
	body := responseBoolean{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// UpdateSwimlane https://docs.kanboard.org/en/latest/api/swimlane_procedures.html#updateswimlane
func (c *Client) UpdateSwimlane(params SwimlaneParams) (bool, error) {
	r := request{
		JSONRPC: "2.0",
		Method:  "getSwimlaneByName",
		ID:      1,
		Params:  params,
	}

	rsp, err := c.Request(r)
	body := responseBoolean{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// AddSwimlane https://docs.kanboard.org/en/latest/api/swimlane_procedures.html#addswimlane
func (c *Client) AddSwimlane(params SwimlaneParams) (int, error) {
	r := request{
		JSONRPC: "2.0",
		Method:  "AddSwimlane",
		ID:      1,
		Params:  params,
	}

	rsp, err := c.Request(r)
	body := responseInt{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// RemoveSwimlane https://docs.kanboard.org/en/latest/api/swimlane_procedures.html#removeswimlane
func (c *Client) RemoveSwimlane(projectID int, swimlaneID int) (bool, error) {
	r := request{
		JSONRPC: "2.0",
		Method:  "RemoveSwimlane",
		ID:      1,
		Params: map[string]int{
			"project_id":  projectID,
			"swimlane_id": swimlaneID,
		},
	}

	rsp, err := c.Request(r)
	body := responseBoolean{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// DisableSwimlane https://docs.kanboard.org/en/latest/api/swimlane_procedures.html#disableswimlane
func (c *Client) DisableSwimlane(projectID int, swimlaneID int) (bool, error) {
	r := request{
		JSONRPC: "2.0",
		Method:  "DisableSwimlane",
		ID:      1,
		Params: map[string]int{
			"project_id":  projectID,
			"swimlane_id": swimlaneID,
		},
	}

	rsp, err := c.Request(r)
	body := responseBoolean{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// EnableSwimlane https://docs.kanboard.org/en/latest/api/swimlane_procedures.html#enableswimlane
func (c *Client) EnableSwimlane(projectID int, swimlaneID int) (bool, error) {
	r := request{
		JSONRPC: "2.0",
		Method:  "EnableSwimlane",
		ID:      1,
		Params: map[string]int{
			"project_id":  projectID,
			"swimlane_id": swimlaneID,
		},
	}

	rsp, err := c.Request(r)
	body := responseBoolean{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// Swimlane type
type Swimlane struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Position  int    `json:"position,string"`
	IsActive  int    `json:"is_active,string"`
	ProjectID int    `json:"project_id,string"`
}

// SwimlaneParams input
type SwimlaneParams struct {
	Description string `json:"description,omitempty"`
	Name        string `json:"name"`
	ProjectID   int    `json:"project_id,string"`
	SwimlaneID  int    `json:"swimlane_id,string,omitempty"`
}

type responseSwimlanes struct {
	JSONRPC string     `json:"jsonrpc"`
	ID      int        `json:"id"`
	Result  []Swimlane `json:"result"`
}

type responseSwimlane struct {
	JSONRPC string   `json:"jsonrpc"`
	ID      int      `json:"id"`
	Result  Swimlane `json:"result"`
}
