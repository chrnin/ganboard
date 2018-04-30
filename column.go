package ganboard

import "encoding/json"

// GetColumns https://docs.kanboard.org/en/latest/api/column_procedures.html#getcolumns
func (c *Client) GetColumns(projectID int) ([]Column, error) {
	r := request{
		JSONRPC: "2.0",
		Method:  "getColumns",
		ID:      1,
		Params: map[string]int{
			"project_id": projectID,
		},
	}

	rsp, err := c.Request(r)
	body := new(responseColumns)
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// GetColumn https://docs.kanboard.org/en/latest/api/column_procedures.html#getcolumn
func (c *Client) GetColumn(columnID int) (Column, error) {
	r := request{
		JSONRPC: "2.0",
		Method:  "getColumn",
		ID:      1,
		Params: map[string]int{
			"column_id": columnID,
		},
	}

	rsp, err := c.Request(r)
	body := new(responseColumn)
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// ChangeColumnPosition https://docs.kanboard.org/en/latest/api/column_procedures.html#changecolumnposition
func (c *Client) ChangeColumnPosition(projectID int, columnID int, position int) (bool, error) {
	r := request{
		JSONRPC: "2.0",
		Method:  "changeColumnPosition",
		ID:      1,
		Params: map[string]int{
			"project_id": projectID,
			"column_id":  columnID,
			"position":   position,
		},
	}

	rsp, err := c.Request(r)
	body := new(responseBoolean)
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// UpdateColumn https://docs.kanboard.org/en/latest/api/column_procedures.html#updatecolumn
func (c *Client) UpdateColumn(params ColumnParams) (bool, error) {
	r := request{
		JSONRPC: "2.0",
		Method:  "updateColumn",
		ID:      1,
		Params:  params,
	}

	rsp, err := c.Request(r)
	body := new(responseBoolean)
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// AddColumn https://docs.kanboard.org/en/latest/api/column_procedures.html#addcolumn
func (c *Client) AddColumn(params ColumnParams) (int, error) {
	r := request{
		JSONRPC: "2.0",
		Method:  "addColumn",
		ID:      1,
		Params:  params,
	}

	rsp, err := c.Request(r)
	body := new(responseInt)
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

// RemoveColumn https://docs.kanboard.org/en/latest/api/column_procedures.html#removecolumn
func (c *Client) RemoveColumn(columnID int) (bool, error) {
	r := request{
		JSONRPC: "2.0",
		Method:  "addColumn",
		ID:      1,
		Params: map[string]int{
			"column_id": columnID,
		},
	}

	rsp, err := c.Request(r)
	body := new(responseBoolean)
	err = json.NewDecoder(rsp.Body).Decode(&body)

	return body.Result, err
}

type responseColumns struct {
	JSONRPC string   `json:"jsonrpc"`
	ID      int      `json:"id"`
	Result  []Column `json:"result"`
}

type responseColumn struct {
	JSONRPC string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  Column `json:"result"`
}

// ColumnParams input for UpdateColumn
type ColumnParams struct {
	ColumnID    int    `json:"column_id"`
	Title       string `json:"title"`
	TaskLimit   int    `json:"task_limit,omitempty"`
	Description string `json:"description,omitempty"`
}
