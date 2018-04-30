package ganboard

import "encoding/json"

// GetAvailableActions https://docs.kanboard.org/en/latest/api/action_procedures.html#getavailableactions
func (c *Client) GetAvailableActions() (map[string]string, error) {
	r := request{
		JSONRPC: "2.0",
		Method:  "getAvailableActions",
		ID:      1,
	}

	rsp, err := c.Request(r)
	body := responseMapStringString{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// GetAvailableActionEvents https://docs.kanboard.org/en/latest/api/action_procedures.html#getavailableactionevents
func (c *Client) GetAvailableActionEvents() (map[string]string, error) {
	r := request{
		JSONRPC: "2.0",
		Method:  "getAvailableActionEvents",
		ID:      1,
	}

	rsp, err := c.Request(r)
	body := responseMapStringString{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// GetCompatibleActionEvents https://docs.kanboard.org/en/latest/api/action_procedures.html#getcompatibleactionevents
func (c *Client) GetCompatibleActionEvents(actionName string) (map[string]string, error) {
	r := request{
		JSONRPC: "2.0",
		Method:  "getCompatibleActionEvents",
		ID:      1,
		Params: map[string]string{
			"action_name": actionName,
		},
	}

	rsp, err := c.Request(r)
	body := responseMapStringString{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// GetActions https://docs.kanboard.org/en/latest/api/action_procedures.html#getactions
func (c *Client) GetActions(projectID int) (Action, error) {
	r := request{
		JSONRPC: "2.0",
		Method:  "getActions",
		ID:      1,
		Params: map[string]int{
			"project_id": projectID,
		},
	}

	rsp, err := c.Request(r)
	body := responseAction{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// CreateAction https://docs.kanboard.org/en/latest/api/action_procedures.html#createaction
func (c *Client) CreateAction(params ActionParams) (int, error) {
	r := request{
		JSONRPC: "2.0",
		Method:  "createAction",
		ID:      1,
		Params:  params,
	}

	rsp, err := c.Request(r)
	body := responseInt{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// RemoveAction https://docs.kanboard.org/en/latest/api/action_procedures.html#createaction
func (c *Client) RemoveAction(actionID int) (bool, error) {
	r := request{
		JSONRPC: "2.0",
		Method:  "removeAction",
		ID:      1,
		Params: map[string]int{
			"action_id": actionID,
		},
	}

	rsp, err := c.Request(r)
	body := responseBoolean{}
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// Action type
type Action struct {
	ID         int               `json:"id,string"`
	ProjectID  int               `json:"project_id,string"`
	EventName  string            `json:"event_name"`
	ActionName string            `json:"action_name"`
	Params     map[string]string `json:"params"`
}

// ActionParams input for CreateAction
type ActionParams struct {
	ProjectID  int               `json:"project_id,string"`
	EventName  string            `json:"event_name"`
	ActionName string            `json:"action_name"`
	Params     map[string]string `json:"params"`
}

type responseAction struct {
	JSONRPC string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  Action `json:"result"`
}
