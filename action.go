package ganboard

import "encoding/json"

// GetAvailableActions https://docs.kanboard.org/en/latest/api/action_procedures.html#getavailableactions
func (c *Client) GetAvailableActions() (map[string]string, error) {
	query := request{
		Client: c,
		Method: "getAvailableActions",
		ID:     1,
	}
	response, err := query.decodeMapStringString()
	return response, err
}

// GetAvailableActionEvents https://docs.kanboard.org/en/latest/api/action_procedures.html#getavailableactionevents
func (c *Client) GetAvailableActionEvents() (map[string]string, error) {
	query := request{
		Client: c,
		Method: "getAvailableActionEvents",
		ID:     1,
	}
	response, err := query.decodeMapStringString()
	return response, err
}

// GetCompatibleActionEvents https://docs.kanboard.org/en/latest/api/action_procedures.html#getcompatibleactionevents
func (c *Client) GetCompatibleActionEvents(actionName string) (map[string]string, error) {
	query := request{
		Client: c,
		Method: "getCompatibleActionEvents",
		Params: map[string]string{
			"action_name": actionName,
		},
	}
	response, err := query.decodeMapStringString()
	return response, err
}

// GetActions https://docs.kanboard.org/en/latest/api/action_procedures.html#getactions
func (c *Client) GetActions(projectID int) (Action, error) {
	query := request{
		Client: c,
		Method: "getActions",
		Params: map[string]int{
			"project_id": projectID,
		},
	}
	response, err := query.decodeAction()
	return response, err
}

// CreateAction https://docs.kanboard.org/en/latest/api/action_procedures.html#createaction
func (c *Client) CreateAction(params ActionParams) (int, error) {
	query := request{
		Client: c,
		Method: "createAction",
		Params: params,
	}
	response, err := query.decodeInt()
	return response, err
}

// RemoveAction https://docs.kanboard.org/en/latest/api/action_procedures.html#createaction
func (c *Client) RemoveAction(actionID int) (bool, error) {
	query := request{
		Client: c,
		Method: "removeAction",
		Params: map[string]int{
			"action_id": actionID,
		},
	}
	response, err := query.decodeBoolean()
	return response, err
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

func (r *request) decodeAction() (Action, error) {
	rsp, err := r.Client.Request(*r)
	if err != nil {
		return Action{}, err
	}

	body := struct {
		JSONRPC string `json:"jsonrpc"`
		ID      int    `json:"id"`
		Result  Action `json:"result"`
	}{}

	err = json.NewDecoder(rsp.Body).Decode(&body)
	return body.Result, err
}
