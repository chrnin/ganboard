package ganboard

import "encoding/json"

// GetActiveSwimlanes https://docs.kanboard.org/en/latest/api/swimlane_procedures.html#swimlane-api-procedures
func (c *Client) GetActiveSwimlanes(projectID int) ([]Swimlane, error) {
	query := request{
		Client: c,
		Method: "getActiveSwimlanes",
		Params: map[string]int{
			"project_id": projectID,
		},
	}
	response, err := query.decodeSwimlanes()
	return response, err
}

// GetAllSwimlanes https://docs.kanboard.org/en/latest/api/swimlane_procedures.html#getallswimlanes
func (c *Client) GetAllSwimlanes(projectID int) ([]Swimlane, error) {
	query := request{
		Client: c,
		Method: "getAllSwimlanes",
		Params: map[string]int{
			"project_id": projectID,
		},
	}
	response, err := query.decodeSwimlanes()
	return response, err
}

// GetSwimlane https://docs.kanboard.org/en/latest/api/swimlane_procedures.html#getswimlane
func (c *Client) GetSwimlane(swimlaneID int) (Swimlane, error) {
	query := request{
		Client: c,
		Method: "getSwimlane",
		Params: map[string]int{
			"swimlane_id": swimlaneID,
		},
	}
	response, err := query.decodeSwimlane()
	return response, err
}

// GetSwimlaneByID https://docs.kanboard.org/en/latest/api/swimlane_procedures.html#getswimlanebyid
func (c *Client) GetSwimlaneByID(swimlaneID int) (Swimlane, error) {
	query := request{
		Client: c,
		Method: "getSwimlaneById",
		Params: map[string]int{
			"swimlane_id": swimlaneID,
		},
	}
	response, err := query.decodeSwimlane()
	return response, err
}

// GetSwimlaneByName https://docs.kanboard.org/en/latest/api/swimlane_procedures.html#getswimlanebyname
func (c *Client) GetSwimlaneByName(projectID int, name string) (Swimlane, error) {
	query := request{
		Client: c,
		Method: "getSwimlaneByName",
		Params: struct {
			ProjectID int    `json:"project_id"`
			Name      string `json:"Name"`
		}{
			ProjectID: projectID,
			Name:      name,
		},
	}
	response, err := query.decodeSwimlane()
	return response, err
}

// ChangeSwimlanePosition https://docs.kanboard.org/en/latest/api/swimlane_procedures.html#changeswimlaneposition
func (c *Client) ChangeSwimlanePosition(projectID int, swimlaneID int, position int) (bool, error) {
	query := request{
		Client: c,
		Method: "getSwimlaneByName",
		Params: map[string]int{
			"project_id":  projectID,
			"swimlane_id": swimlaneID,
			"position":    position,
		},
	}
	response, err := query.decodeBoolean()
	return response, err
}

// UpdateSwimlane https://docs.kanboard.org/en/latest/api/swimlane_procedures.html#updateswimlane
func (c *Client) UpdateSwimlane(params SwimlaneParams) (bool, error) {
	query := request{
		Client: c,
		Method: "getSwimlaneByName",
		Params: params,
	}
	response, err := query.decodeBoolean()
	return response, err
}

// AddSwimlane https://docs.kanboard.org/en/latest/api/swimlane_procedures.html#addswimlane
func (c *Client) AddSwimlane(params SwimlaneParams) (int, error) {
	query := request{
		Client: c,
		Method: "AddSwimlane",
		Params: params,
	}
	response, err := query.decodeInt()
	return response, err
}

// RemoveSwimlane https://docs.kanboard.org/en/latest/api/swimlane_procedures.html#removeswimlane
func (c *Client) RemoveSwimlane(projectID int, swimlaneID int) (bool, error) {
	query := request{
		Client: c,
		Method: "RemoveSwimlane",
		Params: map[string]int{
			"project_id":  projectID,
			"swimlane_id": swimlaneID,
		},
	}
	response, err := query.decodeBoolean()
	return response, err
}

// DisableSwimlane https://docs.kanboard.org/en/latest/api/swimlane_procedures.html#disableswimlane
func (c *Client) DisableSwimlane(projectID int, swimlaneID int) (bool, error) {
	query := request{
		Client: c,
		Method: "DisableSwimlane",
		Params: map[string]int{
			"project_id":  projectID,
			"swimlane_id": swimlaneID,
		},
	}
	response, err := query.decodeBoolean()
	return response, err
}

// EnableSwimlane https://docs.kanboard.org/en/latest/api/swimlane_procedures.html#enableswimlane
func (c *Client) EnableSwimlane(projectID int, swimlaneID int) (bool, error) {
	query := request{
		Client: c,
		Method: "EnableSwimlane",
		Params: map[string]int{
			"project_id":  projectID,
			"swimlane_id": swimlaneID,
		},
	}
	response, err := query.decodeBoolean()
	return response, err
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

func (r *request) decodeSwimlanes() ([]Swimlane, error) {
	rsp, err := r.Client.Request(*r)
	if err != nil {
		return nil, err
	}

	body := struct {
		JSONRPC string     `json:"jsonrpc"`
		ID      int        `json:"id"`
		Result  []Swimlane `json:"result"`
	}{}

	err = json.NewDecoder(rsp.Body).Decode(&body)
	return body.Result, err
}

func (r *request) decodeSwimlane() (Swimlane, error) {
	rsp, err := r.Client.Request(*r)
	if err != nil {
		return Swimlane{}, err
	}

	body := struct {
		JSONRPC string   `json:"jsonrpc"`
		ID      int      `json:"id"`
		Result  Swimlane `json:"result"`
	}{}

	err = json.NewDecoder(rsp.Body).Decode(&body)
	return body.Result, err
}
